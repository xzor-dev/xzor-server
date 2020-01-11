package twitch

import (
	"errors"
	"io"
	"log"

	"github.com/xzor-dev/xzor-server/svc/platform"
	"github.com/xzor-dev/xzor-server/svc/platform/twitch/api"
)

// Game implements platform.Game for twitch games.
type Game struct {
	name   string
	params platform.GameParams
}

// Name returns the name of the game according to twitch.
func (g *Game) Name() string {
	return g.name
}

// Params returns all game parameters returned from twitch.
func (g *Game) Params() platform.GameParams {
	return g.params
}

// GameCollector collects top games from twitch.
type GameCollector struct {
	API        *api.API
	MaxResults int

	game     chan *api.Game
	done     chan bool
	err      chan error
	started  bool
	index    int
	maxIndex int
}

// NewGameCollector creates a new GameCollector instance.
func NewGameCollector(a *api.API, maxResults int) *GameCollector {
	return &GameCollector{
		API:        a,
		MaxResults: maxResults,
	}
}

// Next will return the next game until the maximum number of games is reached
// or no more games can be found.
func (c *GameCollector) Next() (platform.Game, error) {
	if !c.started {
		return nil, errors.New("collector hasn't been started")
	}
	for {
		select {
		case g := <-c.game:
			return &Game{
				name:   g.Name,
				params: g.Params(),
			}, nil
		case err := <-c.err:
			return nil, err
		case <-c.done:
			return nil, io.EOF
		}
	}
}

// Start initialized all required variables and begins collecting games.
func (c *GameCollector) Start() {
	log.Printf("starting collector")
	c.started = true
	c.index = 0
	c.maxIndex = c.MaxResults - 1
	c.game = make(chan *api.Game)
	c.done = make(chan bool)
	c.err = make(chan error)

	go c.collect()
}

func (c *GameCollector) collect() {
	defer func() {
		c.done <- true
	}()

	var after string
	for {
		res, err := c.API.TopGames(&api.TopGamesParams{
			After: after,
		})
		if err != nil {
			c.err <- err
			return
		}

		after = res.Pagination.Cursor

		for _, g := range res.Games {
			c.game <- g
			c.index++

			if c.index > c.maxIndex {
				return
			}
		}
	}
}

package api

import (
	"encoding/json"
	"errors"
	"github.com/xzor-dev/xzor-server/svc/platform"
)

// ErrEmptyGamesResponse indicates that a request for a game returned zero results.
var ErrEmptyGamesResponse = errors.New("no games found")

// Game holds data for games returned by the twitch API.
type Game struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	BoxArtURL string `json:"box_art_url"`
}

// Params returns the platform.GameParams for the game.
func (g *Game) Params() platform.GameParams {
	return platform.GameParams{
		"id":   g.ID,
		"name": g.Name,
	}
}

// GamesParams are parameters used when requesting games.
type GamesParams struct {
	ID   string
	Name string
}

// ToQueryParams converts the GamesParams to a map used by the API request.
func (p *GamesParams) ToQueryParams() map[string]string {
	params := make(map[string]string)
	if p.ID != "" {
		params["id"] = p.ID
	}
	if p.Name != "" {
		params["name"] = p.Name
	}
	return params
}

// GamesResponse is the response returned when getting games from the API.
type GamesResponse struct {
	Games []*Game `json:"data"`
}

// TopGamesParams are parameters used when requesting the top games.
type TopGamesParams struct {
	After  string
	Before string
	First  int
}

// ToQueryParams converts the parameters to a map used by the API request.
func (p *TopGamesParams) ToQueryParams() map[string]string {
	params := make(map[string]string)
	if p.After != "" {
		params["after"] = p.After
	}
	if p.Before != "" {
		params["before"] = p.Before
	}
	if p.First > 100 {
		p.First = 100
	}
	if p.First > 0 {
		params["first"] = string(p.First)
	}
	return params
}

// TopGamesResponse is the response when querying for top games.
type TopGamesResponse struct {
	Games      []*Game     `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

// Game gets a single game matching the supplied parameters.
func (a *API) Game(params *GamesParams) (*Game, error) {
	p := make(map[string]string)
	if params != nil {
		p = params.ToQueryParams()
	}
	data, err := a.Get("games", p)
	if err != nil {
		return nil, err
	}

	res := &GamesResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	if len(res.Games) == 0 {
		return nil, ErrEmptyGamesResponse
	}

	return res.Games[0], nil
}

// TopGames returns a list of the top games from twitch.
func (a *API) TopGames(params *TopGamesParams) (*TopGamesResponse, error) {
	p := make(map[string]string)
	if params != nil {
		p = params.ToQueryParams()
	}
	data, err := a.Get("games/top", p)
	if err != nil {
		return nil, err
	}

	res := &TopGamesResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

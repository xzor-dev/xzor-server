package memory

import (
	"regexp"
	"sync"

	"github.com/xzor-dev/xzor-server/svc/game"
)

// Storage implements game.Storage for in-memory storage.
type Storage struct {
	games map[game.ID]*game.Game
	mux   sync.Mutex
}

// NewStorage creates a new instance of Storage.
func NewStorage() *Storage {
	return &Storage{
		games: make(map[game.ID]*game.Game),
	}
}

// DeleteGame removes a game from memory.
func (s *Storage) DeleteGame(id game.ID) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.games, id)
	return nil
}

// FindGames returns a list of games matching the query string.
func (s *Storage) FindGames(q string) ([]*game.Game, error) {
	games := make([]*game.Game, 0)
	m := regexp.MustCompile(`^(?i)` + q)
	for _, g := range s.games {
		if m.MatchString(g.Name) {
			games = append(games, g)
		}
	}
	return games, nil
}

// Game returns a single game by its ID.
func (s *Storage) Game(id game.ID) (*game.Game, error) {
	if s.games[id] == nil {
		return nil, game.ErrGameNotFound
	}
	return s.games[id], nil
}

// Games returns all games storaged in memory.
func (s *Storage) Games() ([]*game.Game, error) {
	games := make([]*game.Game, len(s.games))
	i := 0
	for _, g := range s.games {
		games[i] = g
		i++
	}
	return games, nil
}

// SaveGame sets or replaces a game instance in memory.
func (s *Storage) SaveGame(g *game.Game) error {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.games[g.ID] = g
	return nil
}

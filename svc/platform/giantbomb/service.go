package giantbomb

import (
	"github.com/xzor-dev/xzor-server/svc/platform"
	"github.com/xzor-dev/xzor-server/svc/platform/giantbomb/api"
)

const (
	// ID is the ID of the GiantBomb platform.
	ID = "giantbomb"
)

// Service handles interacting with the GiantBomb API.
type Service struct {
	api *api.API
}

// NewService creates a new Service instance with the supplied API.
func NewService(api *api.API) *Service {
	return &Service{
		api: api,
	}
}

// ID returns the ID of the GiantBomb platform.
func (s *Service) ID() platform.ID {
	return ID
}

// Name returns the name of the GiantBomb platform.
func (s *Service) Name() string {
	return "GiantBomb"
}

// FindGames returns a list of games found from searching GiantBomb's API.
func (s *Service) FindGames(q string) ([]platform.Game, error) {
	results, err := s.api.SearchGames(q)
	if err != nil {
		return nil, err
	}

	games := make([]platform.Game, len(results))
	for i, res := range results {
		g := &Game{
			name: res.Name,
			params: platform.GameParams{
				"guid": res.GUID,
			},
		}
		games[i] = g
	}
	return games, nil
}

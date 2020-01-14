package twitch

import (
	"github.com/xzor-dev/xzor-server/svc/platform"
	"github.com/xzor-dev/xzor-server/svc/platform/twitch/api"
)

const (
	// ID is the ID of the twitch platform.
	ID = "twitch"
)

// Service implements platform.Service for twitch.tv
type Service struct {
	api api.API
}

// ID returns the ID of the twitch platform service.
func (s *Service) ID() platform.ID {
	return "twitch"
}

// Name returns the name of the twitch platform service.
func (s *Service) Name() string {
	return "Twitch"
}

// Game returns a single game based on the supplied params.
func (s *Service) Game(params platform.GameParams) (platform.Game, error) {
	gp := &api.GamesParams{}
	if params["id"] != "" {
		gp.ID = params["id"]
	}
	if params["name"] != "" {
		gp.Name = params["name"]
	}
	g, err := s.api.Game(gp)
	if err != nil {
		return nil, err
	}

	return &Game{
		name:   g.Name,
		params: g.Params(),
	}, nil
}

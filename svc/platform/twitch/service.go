package twitch

import "github.com/xzor-dev/xzor-server/svc/platform"

// Service implements platform.Service for twitch.tv
type Service struct{}

// ID returns the ID of the twitch platform service.
func (s *Service) ID() platform.ID {
	return "twitch"
}

// Name returns the name of the twitch platform service.
func (s *Service) Name() string {
	return "Twitch"
}

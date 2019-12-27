package team

import "errors"

// Service exposes all methods used to interact with teams.
type Service struct{}

// NewService creates a new team service instance.
func NewService() *Service {
	return &Service{}
}

// CreateTeam creates and returns a new team from the supplied params.
func (s *Service) CreateTeam(p *CreateTeamParams) (*Team, error) {
	return nil, errors.New("not implemented")
}

// Team returns a team by its ID.
func (s *Service) Team(id ID) (*Team, error) {
	return nil, errors.New("not implemented")
}

// Teams returns a list of teams.
func (s *Service) Teams() ([]*Team, error) {
	return make([]*Team, 0), errors.New("not implemented")
}

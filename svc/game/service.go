package game

import "github.com/xzor-dev/xzor-server/svc/platform"

import "regexp"

// Service handles various functions related to games.
type Service struct {
	platformFinders map[platform.ID]platform.GameFinder
	storage         Storage
}

// NewService creates a new instance of the game service.
func NewService(storage Storage) *Service {
	return &Service{
		platformFinders: make(map[platform.ID]platform.GameFinder),
		storage:         storage,
	}
}

// Game returns a game by its ID.
func (s *Service) Game(id ID) (*Game, error) {
	return s.storage.Game(id)
}

// NewGame creates a new game with the supplied name and saves it to the storage.
func (s *Service) NewGame(name string) (*Game, error) {
	g := NewGame(name, nil)
	err := s.SetGame(g)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// RemoveGame deletes a game from the storage.
func (s *Service) RemoveGame(id ID) error {
	return s.storage.DeleteGame(id)
}

// Search returns a list of games matching the supplied query string.
func (s *Service) Search(q string) ([]*Game, error) {
	return s.storage.FindGames(q)
}

// SearchPlatforms returns a list of unique games found from all registered game finders.
func (s *Service) SearchPlatforms(q string) ([]*PlatformGame, error) {
	games := make([]*PlatformGame, 0)
	for pid, finder := range s.platformFinders {
		pgames, err := finder.FindGames(q)
		if err != nil {
			return nil, err
		}
		for _, pg := range pgames {
			found := false
			m := regexp.MustCompile(`^(?i)` + pg.Name() + `$`)
			for _, g := range games {
				if m.MatchString(g.Name) {
					found = true
					break
				}
			}
			if !found {
				g := &PlatformGame{
					PlatformID: pid,
					Name:       pg.Name(),
					Params:     pg.Params(),
				}
				games = append(games, g)
			}
		}
	}

	return games, nil
}

// SetGame saves a game to the storage.
func (s *Service) SetGame(g *Game) error {
	return s.storage.SaveGame(g)
}

// SetGameFinder sets or replaces a game finder for the supplied platform ID.
func (s *Service) SetGameFinder(platformID platform.ID, finder platform.GameFinder) {
	s.platformFinders[platformID] = finder
}

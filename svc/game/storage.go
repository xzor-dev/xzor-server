package game

// Storage defines functions needed for game storage implementations.
type Storage interface {
	DeleteGame(ID) error
	FindGames(string) ([]*Game, error)
	Game(ID) (*Game, error)
	Games() ([]*Game, error)
	SaveGame(*Game) error
}

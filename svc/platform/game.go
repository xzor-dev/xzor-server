package platform

// Game is an interface for platform-specific games.
type Game interface {
	Name() string
	Params() GameParams
}

// GameCollector is used to collect games from platforms.
type GameCollector interface {
	Next() (Game, error)
}

// GameFinder allows for finding games on a platform.
type GameFinder interface {
	FindGames(string) ([]Game, error)
}

// GameParams are platform-specific parameters for games.
type GameParams map[string]string

// GamePlatform is implemented by platforms that support games.
type GamePlatform interface {
	Game(*GameParams) (Game, error)
}

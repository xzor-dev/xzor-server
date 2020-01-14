package giantbomb

import "github.com/xzor-dev/xzor-server/svc/platform"

// Game holds data for a game from GiantBomb.
type Game struct {
	name   string
	params platform.GameParams
}

// Name returns the name of the game.
func (g *Game) Name() string {
	return g.name
}

// Params returns the game's platform-specific parameters.
func (g *Game) Params() platform.GameParams {
	return g.params
}

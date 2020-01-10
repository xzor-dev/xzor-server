package game

import "github.com/xzor-dev/xzor-server/svc/platform"

// Game holds information for a single game.
type Game struct {
	ID             ID                                  `json:"id"`
	Name           string                              `json:"name"`
	PlatformParams map[platform.ID]platform.GameParams `json:"platformParams"`
}

// ID is a unique string for a game.
type ID string

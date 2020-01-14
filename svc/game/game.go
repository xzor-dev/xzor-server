package game

import (
	"errors"

	"github.com/xzor-dev/xzor-server/lib"
	"github.com/xzor-dev/xzor-server/svc/platform"
)

const (
	// IDLength is the default length of generated game IDs.
	IDLength = 10
)

var (
	// ErrGameNotFound indicates that a game could not be found.
	ErrGameNotFound = errors.New("game not found")
)

// Game holds information for a single game.
type Game struct {
	ID             ID                                  `json:"id"`
	Name           string                              `json:"name"`
	PlatformParams map[platform.ID]platform.GameParams `json:"platformParams"`
}

// NewGame creates a new game instance with the supplied name and an optional ID.
// If an ID is not provided, a new ID will be generated.
func NewGame(name string, id *ID) *Game {
	if id == nil {
		gid := NewID()
		id = &gid
	}
	return &Game{
		ID:             *id,
		Name:           name,
		PlatformParams: make(map[platform.ID]platform.GameParams),
	}
}

// ID is a unique string for a game.
type ID string

// NewID generates a new game ID.
func NewID() ID {
	id := lib.NewRandomString(10)
	return ID(id)
}

// PlatformGame holds data for a game from a platform.
type PlatformGame struct {
	PlatformID platform.ID         `json:"platformID"`
	Name       string              `json:"name"`
	Params     platform.GameParams `json:"params"`
}

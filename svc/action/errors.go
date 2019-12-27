package action

import "errors"

// ErrInvalidSourceName indicates that a source name is not registered.
var ErrInvalidSourceName = errors.New("invalid action source name")

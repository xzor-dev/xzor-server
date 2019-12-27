package chain

import "errors"

// ErrInvalidBlockIndex indicates that a block's index is invalid.
var ErrInvalidBlockIndex = errors.New("invalid block index")

// ErrInvalidPreviousBlockHash indicates that a block's previous block hash is invalid.
var ErrInvalidPreviousBlockHash = errors.New("invalid previous block hash")

// ErrNoPendingActions indicates that the chain service has no pending actions.
var ErrNoPendingActions = errors.New("no pending actions available")

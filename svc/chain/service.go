package chain

import (
	"encoding/json"
	"time"

	"github.com/xzor-dev/xzor-server/svc/action"
)

// Service controls reading and writing from the chain.
type Service struct {
	Chain *Chain

	pendingActions []*action.Action
}

// NewService creates a new chain service for the supplied chain.
func NewService(c *Chain) *Service {
	return &Service{
		Chain: c,
	}
}

// AddAction adds an action to a list of pending actions.
func (s *Service) AddAction(a *action.Action) {
	s.pendingActions = append(s.pendingActions, a)
}

// NewBlock generates a new block using the pending actions in the service.
// After successful creation, the block is added to the chain and all pending
// actions are cleared.
func (s *Service) NewBlock() (*Block, error) {
	if (len(s.pendingActions)) == 0 {
		return nil, ErrNoPendingActions
	}
	data, err := json.Marshal(s.pendingActions)
	if err != nil {
		return nil, err
	}
	t := time.Now()
	prevBlock := s.Chain.LastBlock
	block, err := NewBlock(data, t.Unix(), prevBlock)
	if err != nil {
		return nil, err
	}
	err = s.Chain.Push(block)
	if err != nil {
		return nil, err
	}
	s.pendingActions = make([]*action.Action, 0)
	return block, nil
}

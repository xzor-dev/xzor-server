package chain

// Chain holds an ordered set of blocks.
type Chain struct {
	Blocks    map[BlockHash]*Block
	LastBlock *Block
}

// New creates an empty chain.
func New() *Chain {
	return &Chain{
		Blocks: make(map[BlockHash]*Block),
	}
}

// Push adds a block to the chain while enforcing its ordering.
func (c *Chain) Push(b *Block) error {
	if c.Blocks[b.Hash] != nil {
		return nil
	}
	if b.PreviousHash == "" && len(c.Blocks) != 0 {
		return ErrInvalidPreviousBlockHash
	}
	if b.PreviousHash != "" {
		if len(c.Blocks) == 0 {
			return ErrInvalidPreviousBlockHash
		}
		if prevBlock := c.Blocks[b.PreviousHash]; prevBlock == nil {
			return ErrInvalidPreviousBlockHash
		} else if prevBlock.Index != b.Index-1 {
			return ErrInvalidBlockIndex
		}
	}

	c.Blocks[b.Hash] = b
	c.LastBlock = b
	return nil
}

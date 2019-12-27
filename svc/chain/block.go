package chain

import (
	"github.com/xzor-dev/xzor-server/lib"
)

// BlockHash is a unique string for a block.
type BlockHash string

// NewBlockHash generates a unique block hash.
func NewBlockHash(prevHash BlockHash, index int, data []byte, timestamp int64) (BlockHash, error) {
	rs := string(index) + "--" + string(timestamp) + "--" + string(prevHash) + "--"
	rb := []byte(rs)
	rb = append(rb, data...)
	hash, err := lib.NewHash(rb)
	if err != nil {
		return "", err
	}
	return BlockHash(hash), nil
}

// Block holds multiple actions.
type Block struct {
	Hash         BlockHash
	PreviousHash BlockHash
	Index        int
	Data         []byte
	Timestamp    int64
}

// NewBlock creates a new block at the next index.
// If no previous block is provided the block is treated as the genesis block.
func NewBlock(data []byte, timestamp int64, prevBlock *Block) (*Block, error) {
	var index int
	var prevHash BlockHash
	if prevBlock != nil {
		index = prevBlock.Index + 1
		prevHash = prevBlock.Hash
	}

	hash, err := NewBlockHash(prevHash, index, data, timestamp)
	if err != nil {
		return nil, err
	}
	return &Block{
		PreviousHash: prevHash,
		Hash:         hash,
		Index:        index,
		Data:         data,
		Timestamp:    timestamp,
	}, nil
}

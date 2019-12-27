package action

import (
	"time"

	"github.com/xzor-dev/xzor-server/lib"
	"github.com/xzor-dev/xzor-server/svc/client"
)

// Hash is a unique string assigned to actions.
type Hash string

// NewHash generates a unique action hash.
func NewHash(cli *client.Client, sourceName SourceName, actionName string, data []byte, timestamp int64) (Hash, error) {
	ch, err := cli.Hash()
	if err != nil {
		return "", err
	}
	rs := string(ch) + "--" + string(sourceName) + "--" + actionName + "--" + string(timestamp)
	rb := []byte(rs)
	rb = append(rb, data...)
	hash, err := lib.NewHash(rb)
	if err != nil {
		return "", err
	}
	return Hash(hash), nil
}

// Action holds data of a single action.
type Action struct {
	Hash      Hash
	Name      string
	Source    SourceName
	Data      []byte
	Timestamp int64
}

// New creates a new action.
func New(cli *client.Client, sourceName SourceName, actionName string, data []byte) (*Action, error) {
	t := time.Now()
	hash, err := NewHash(cli, sourceName, actionName, data, t.Unix())
	if err != nil {
		return nil, err
	}
	return &Action{
		Hash:      hash,
		Name:      actionName,
		Source:    sourceName,
		Data:      data,
		Timestamp: t.Unix(),
	}, nil
}

package chain_test

import (
	"testing"

	"github.com/xzor-dev/xzor-server/svc/action"
	"github.com/xzor-dev/xzor-server/svc/chain"
	"github.com/xzor-dev/xzor-server/svc/user"
)

func TestChainService(t *testing.T) {
	c := chain.New()
	u, err := user.New()
	if err != nil {
		t.Fatalf("%v", err)
	}
	service := chain.NewService(c)
	_, err = service.NewBlock()
	if err == nil {
		t.Fatal("expected an error when creating a new block")
	}
	act, err := action.New(u, "test-source", "test-action", []byte("test-data"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	service.AddAction(act)
	b, err := service.NewBlock()
	if err != nil {
		t.Fatalf("%v", err)
	}
	if c.LastBlock == nil || c.LastBlock.Hash != b.Hash {
		t.Fatal("expected chain's last block to be newly created block")
	}
}

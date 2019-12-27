package action_test

import (
	"testing"

	"github.com/xzor-dev/xzor-server/svc/action"
	"github.com/xzor-dev/xzor-server/svc/client"
)

func TestActionService(t *testing.T) {
	sourceName := action.SourceName("test-source")
	sourceInstance := &testSource{}
	service := action.NewService()
	service.SetSource(sourceName, sourceInstance)

	cli, err := client.New()
	if err != nil {
		t.Fatalf("%v", err)
	}
	act, err := action.New(cli, sourceName, "test-func", []byte("hello"))
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = service.ProcessAction(act)
	if err != nil {
		t.Fatalf("%v", err)
	}
	testAct := sourceInstance.lastAction
	if testAct == nil {
		t.Fatal("action was not properly routes to test source")
	}
}

type testSource struct {
	lastAction *action.Action
}

func (s *testSource) ProcessAction(a *action.Action) error {
	s.lastAction = a
	return nil
}

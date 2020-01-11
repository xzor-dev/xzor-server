package twitch_test

import (
	"io"
	"os"
	"testing"

	"github.com/xzor-dev/xzor-server/svc/platform"
	"github.com/xzor-dev/xzor-server/svc/platform/twitch"
	"github.com/xzor-dev/xzor-server/svc/platform/twitch/api"
)

func TestGameCollector(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("%v", err)
	}

	a, err := api.FromConfigFile(dir + "/api/testdata/api_config.json")
	if err != nil {
		t.Fatalf("%v", err)
	}

	max := 25
	collector := twitch.NewGameCollector(a, max)
	collector.Start()
	games := make([]platform.Game, max)
	i := 0
	for {
		g, err := collector.Next()
		if err != nil {
			if err != io.EOF {
				t.Fatalf("%v", err)
			}
			break
		}
		games[i] = g
		i++
	}
	if len(games) != max {
		t.Fatalf("unexpected number of games collected: wanted %d, got %d", max, len(games))
	}
}

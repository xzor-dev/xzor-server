package api_test

import (
	"os"
	"testing"

	"github.com/xzor-dev/xzor-server/svc/platform/giantbomb/api"
)

func TestAPI(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("%v", err)
	}

	a, err := api.FromConfigFile(dir + "/testdata/api_config.json")
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Run("Get Single Game", func(t *testing.T) {
		guid := "3030-66622"
		name := "Battlefield V"
		g, err := a.Game(guid)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if g.Name != name {
			t.Fatalf("unexpected game name: wanted %s, got %s", name, g.Name)
		}
	})

	t.Run("Search For Game", func(t *testing.T) {
		guid := "3030-72014"
		name := "Apex Legends"
		games, err := a.SearchGames(name)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if len(games) < 1 {
			t.Fatalf("expected at least 1 result, got 0")
		}
		if games[0].GUID != guid {
			t.Fatalf("expected first result's GUID to be %s, got %s", guid, games[0].GUID)
		}
	})
}

package api_test

import (
	"os"
	"testing"

	"github.com/xzor-dev/xzor-server/svc/platform/twitch/api"
)

func TestGamesAPI(t *testing.T) {
	a, err := newAPI()
	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Run("Get Game By Name", func(t *testing.T) {
		name := "Apex Legends"
		g, err := a.Game(&api.GamesParams{
			Name: name,
		})
		if err != nil {
			t.Fatalf("%v", err)
		}
		if g.Name != name {
			t.Fatalf("unexpected game name: wanted %s, got %s", name, g.Name)
		}
	})

	t.Run("Get Game By ID", func(t *testing.T) {
		id := "504199"
		name := "Battlefield V"
		g, err := a.Game(&api.GamesParams{
			ID: id,
		})
		if err != nil {
			t.Fatalf("%v", err)
		}
		if g.Name != name {
			t.Fatalf("unexpected game name: wanted %s, got %s", name, g.Name)
		}
	})

	t.Run("Get Top Games", func(t *testing.T) {
		res, err := a.TopGames(nil)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if len(res.Games) == 0 {
			t.Fatalf("expected a list of games, got: %v", res)
		}
	})
}

func TestStreamsAPI(t *testing.T) {
	a, err := newAPI()
	if err != nil {
		t.Fatalf("%v", err)
	}

	res, err := a.Streams(&api.StreamsParams{
		GameID: "504199",
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(res.Streams) == 0 {
		t.Fatalf("expected at least one stream, got %v", res)
	}
}

func TestUsersAPI(t *testing.T) {
	a, err := newAPI()
	if err != nil {
		t.Fatalf("%v", err)
	}

	_, err = a.User(&api.UserParams{
		Login: "iainz0r",
	})
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func newAPI() (*api.API, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return api.FromConfigFile(dir + "/testdata/api_config.json")
}

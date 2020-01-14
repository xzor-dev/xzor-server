package game_test

import (
	"os"
	"testing"

	"github.com/xzor-dev/xzor-server/svc/game"
	"github.com/xzor-dev/xzor-server/svc/game/memory"
	"github.com/xzor-dev/xzor-server/svc/platform/giantbomb"
	giantbomb_api "github.com/xzor-dev/xzor-server/svc/platform/giantbomb/api"
)

// Search for game from game service (expect 0 results)
// Search for game from platforms supporting game search (expect at least 1 result)
// Take first result and add it to the system using the game service
// Search for the game again using the game service (expect 1 result)
// Load game details using its ID returned from service search
// Delete game
func TestNewGameFlow(t *testing.T) {
	gbService := newGBService(t)

	name := "Apex Legends"
	storage := memory.NewStorage()
	gs := game.NewService(storage)
	gs.SetGameFinder(giantbomb.ID, gbService)

	games, err := gs.Search(name)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(games) > 0 {
		t.Fatalf("expected 0 results, got %d", len(games))
	}

	pGames, err := gs.SearchPlatforms(name)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(pGames) == 0 {
		t.Fatal("expected at least 1 result, got 0")
	}

	pGame := pGames[0]
	if pGame.Name != name {
		t.Fatalf("unexpected game name: wanted %s, got %s", name, pGame.Name)
	}
	g, err := gs.NewGame(pGame.Name)
	if err != nil {
		t.Fatalf("%v", err)
	}
	g.PlatformParams[pGame.PlatformID] = pGame.Params

	err = gs.SetGame(g)
	if err != nil {
		t.Fatalf("%v", err)
	}
	games, err = gs.Search(name)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(games) != 1 {
		t.Fatalf("expected 1 result, got %d", len(games))
	}
	_, err = gs.Game(g.ID)
	if err != nil {
		t.Fatalf("%v", err)
	}
	err = gs.RemoveGame(g.ID)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func newGBService(t *testing.T) *giantbomb.Service {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("%v", err)
	}
	api, err := giantbomb_api.FromConfigFile(dir + "/testdata/giantbomb_api.json")
	if err != nil {
		t.Fatalf("%v", err)
	}
	return giantbomb.NewService(api)
}

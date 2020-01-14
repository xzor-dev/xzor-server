package api

import (
	"encoding/json"
)

// Game holds data on a single game.
type Game struct {
	Aliases         string `json:"aliases"`
	APIDetailURL    string `json:"api_detail_url"`
	DateAdded       string `json:"date_added"`
	DateLastUpdated string `json:"date_last_updated"`
	Description     string `json:"description"`
	GUID            string `json:"guid"`
	Name            string `json:"name"`
}

// GameResponse is the response from querying a single game.
type GameResponse struct {
	StandardResponse

	Game *Game `json:"results"`
}

// GamesResponse is the response from querying multiple games.
type GamesResponse struct {
	StandardResponse

	Games []*Game `json:"results"`
}

// Game returns a single game by its GUID.
func (a *API) Game(guid string) (*Game, error) {
	data, err := a.Get("game/"+guid, nil)
	if err != nil {
		return nil, err
	}

	res := &GameResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res.Game, nil
}

// SearchGames finds all games matching the supplied query.
func (a *API) SearchGames(query string) ([]*Game, error) {
	data, err := a.Get("search", map[string]string{
		"query":     query,
		"resources": "game",
	})
	if err != nil {
		return nil, err
	}

	res := &GamesResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res.Games, nil
}

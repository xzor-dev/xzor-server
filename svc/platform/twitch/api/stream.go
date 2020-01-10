package api

import "encoding/json"

// Stream holds data for a stream.
type Stream struct {
	GameID       string      `json:"game_id"`
	ID           string      `json:"id"`
	Language     string      `json:"language"`
	Pagination   *Pagination `json:"pagination"`
	StartedAt    string      `json:"started_at"`
	TagIDs       []string    `json:"tag_ids"`
	ThumbnailURL string      `json:"thumbnail_url"`
	Title        string      `json:"title"`
	Type         string      `json:"type"`
	UserID       string      `json:"user_id"`
	Username     string      `json:"user_name"`
	ViewerCount  int         `json:"viewer_count"`
}

// StreamsParams are parameters used when requesting streams.
type StreamsParams struct {
	After     string
	Before    string
	First     int
	GameID    string
	Language  string
	UserID    string
	UserLogin string
}

// ToQueryParams converts available StreamsParams values to a map used when making requests.
func (p *StreamsParams) ToQueryParams() map[string]string {
	params := make(map[string]string)
	if p.After != "" {
		params["after"] = p.After
	}
	if p.Before != "" {
		params["before"] = p.Before
	}
	if p.First > 100 {
		p.First = 100
	}
	if p.First > 0 {
		params["first"] = string(p.First)
	}
	if p.GameID != "" {
		params["game_id"] = p.GameID
	}
	if p.Language != "" {
		params["language"] = p.Language
	}
	if p.UserID != "" {
		params["user_id"] = p.UserID
	}
	if p.UserLogin != "" {
		params["user_login"] = p.UserLogin
	}
	return params
}

// StreamsResponse is the response returned when querying for streams.
type StreamsResponse struct {
	Streams    []*Stream   `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

// Streams returns a list of streams matching the supplied parameters.
func (a *API) Streams(p *StreamsParams) (*StreamsResponse, error) {
	params := make(map[string]string)
	if p != nil {
		params = p.ToQueryParams()
	}
	data, err := a.Get("streams", params)
	if err != nil {
		return nil, err
	}
	res := &StreamsResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

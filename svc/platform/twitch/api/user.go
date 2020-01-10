package api

import (
	"encoding/json"
	"errors"
)

// ErrEmptyUsersResponse indicates that a request for users returned zero results.
var ErrEmptyUsersResponse = errors.New("no users found")

// User holds data for a single user.
type User struct {
	ID              string `json:"id"`
	Email           string `json:"email"`
	DisplayName     string `json:"display_name"`
	Description     string `json:"description"`
	Login           string `json:"login"`
	BroadcasterType string `json:"broadcaster_type"`
	OfflineImageURL string `json:"offline_image_url"`
	ProfileImageURL string `json:"profile_image_url"`
	Type            string `json:"type"`
	ViewCount       int    `json:"view_count"`
}

// UserParams are used when querying for a single user.
type UserParams struct {
	ID    string
	Login string
}

// ToQueryParams converts UserParams to a map used when requesting a user.
func (p *UserParams) ToQueryParams() map[string]string {
	params := make(map[string]string)
	if p.ID != "" {
		params["id"] = p.ID
	}
	if p.Login != "" {
		params["login"] = p.Login
	}
	return params
}

// UsersResponse is the response when querying users.
type UsersResponse struct {
	Users []*User `json:"data"`
}

// User returns a single user based on the supplied parameters.
func (a *API) User(p *UserParams) (*User, error) {
	params := make(map[string]string)
	if p != nil {
		params = p.ToQueryParams()
	}

	data, err := a.Get("users", params)
	if err != nil {
		return nil, err
	}

	res := &UsersResponse{}
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	if len(res.Users) == 0 {
		return nil, ErrEmptyUsersResponse
	}

	return res.Users[0], nil
}

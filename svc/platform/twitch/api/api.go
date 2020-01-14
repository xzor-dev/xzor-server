package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// BaseURL is the URL all requests are made to.
	BaseURL = "https://api.twitch.tv/helix"
)

// Config holds configuration data for the API.
type Config struct {
	ClientID string `json:"clientID"`
}

// API provides methods used to interact with the twitch API.
type API struct {
	ClientID string
}

// FromConfigFile creates a new API instance using a configuration file location at path.
func FromConfigFile(path string) (*API, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	conf := &Config{}
	err = json.Unmarshal(data, conf)
	if err != nil {
		return nil, err
	}

	return &API{
		ClientID: conf.ClientID,
	}, nil
}

// Get executes a GET request on the supplied path and adds the supplied params as query parameters.
func (a *API) Get(path string, params map[string]string) ([]byte, error) {
	pairs := make([]string, 0)
	for key, value := range params {
		pairs = append(pairs, key+"="+url.QueryEscape(value))
	}
	query := strings.Join(pairs, "&")
	url := BaseURL + "/" + path + "?" + query
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Client-ID", a.ClientID)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return data, nil
}

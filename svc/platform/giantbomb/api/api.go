package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	// BaseURL is the URL used for all API requests.
	BaseURL = "https://www.giantbomb.com/api"
)

// API is used for interacting with the GiantBomb API.
type API struct {
	Key string
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
		Key: conf.APIKey,
	}, nil
}

// Get executes a GET request on the supplied path with any addition query parameters supplied in params.
// All requests made with `api_key` set to the predefined API key and `format` set to `json`.
func (a *API) Get(path string, params map[string]string) ([]byte, error) {
	pairs := make([]string, 0)
	if params == nil {
		params = make(map[string]string)
	}
	params["api_key"] = a.Key
	params["format"] = "json"

	for key, value := range params {
		pairs = append(pairs, key+"="+url.QueryEscape(value))
	}
	query := strings.Join(pairs, "&")
	url := BaseURL + "/" + path + "/?" + query
	log.Printf("GET %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

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

// Config holds configuration data for the API.
type Config struct {
	APIKey string `json:"api_key"`
}

// StandardResponse is returned with most API responses.
type StandardResponse struct {
	Error        string `json:"error"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	TotalPages   int    `json:"number_of_page_results"`
	TotalResults int    `json:"number_of_total_results"`
	StatusCode   int    `json:"status_code"`
}

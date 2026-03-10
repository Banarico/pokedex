package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type RespShallowLocations struct {
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

// ListLocations fetches a page of location areas from the PokeAPI.
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		// If a URL was provided (next or previous), use that instead!
		url = *pageURL
	}

	// 1. Create the request
        if val, ok := c.cache.Get(url); ok {
            locationsResp := RespShallowLocations{}
            err := json.Unmarshal(val, &locationsResp)
            if err != nil {
                return RespShallowLocations{}, err
            }
            return locationsResp, nil
        }
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// 2. Do the request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	// 3. Read the body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// 4. Unmarshal into your struct
	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

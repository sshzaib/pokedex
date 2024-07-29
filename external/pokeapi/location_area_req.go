package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/sshzaib/pokedex/external/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

func (c *Client) GetLocationAreas(pageURL *string, cache *pokecache.Cache) (LocationAreasResponse, error) {
	endpoint := "/location-area"

	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	if data, ok := cache.GetCache(fullURL); ok {
		locationAreaResponse := LocationAreasResponse{}
		if err := json.Unmarshal(data, &locationAreaResponse); err != nil {
			return LocationAreasResponse{}, err
		}
		for _, locationArea := range locationAreaResponse.Results {
			fmt.Printf("- %v\n", locationArea.Name)
		}
		return locationAreaResponse, nil
	}
	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	if res.StatusCode > 299 {
		fmt.Printf("Error status code: %v\n", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	locationAreaResponse := LocationAreasResponse{}
	if err := json.Unmarshal(body, &locationAreaResponse); err != nil {
		return LocationAreasResponse{}, err
	}
	for _, locationArea := range locationAreaResponse.Results {
		fmt.Printf("- %v\n", locationArea.Name)
	}
	cache.AddCache(fullURL, body)
	return locationAreaResponse, nil
}

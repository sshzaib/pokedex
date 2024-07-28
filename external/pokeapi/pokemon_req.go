package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/sshzaib/pokedex/external/pokecache"
)

func (c *Client) GetPokemonInLocation(location []string, cache pokecache.Cache) (Pokemon, error) {
	endpoint := "/location-area/" + location[1]
	fullURL := baseURL + endpoint
	if data, ok := cache.GetCache(fullURL); ok {
		fmt.Println("cache hit")
		pokemon := Pokemon{}
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, err
		}
		for _, poke := range pokemon.PokemonEncounters {
			fmt.Printf("- %v\n", poke.Pokemon.Name)
		}
		return Pokemon{}, nil
	}
	fmt.Println("cache miss")
	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		return Pokemon{}, err
	}
	if res.StatusCode > 299 {
		fmt.Printf("Error status code: %v\n", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	for _, poke := range pokemon.PokemonEncounters {
		fmt.Printf("- %v\n", poke.Pokemon.Name)
	}
	cache.AddCache(fullURL, body)
	return pokemon, nil
}

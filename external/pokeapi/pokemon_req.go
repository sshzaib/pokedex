package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand/v2"

	"github.com/sshzaib/pokedex/external/pokecache"
)

func (c *Client) GetPokemonInLocation(location []string, cache *pokecache.Cache) (Pokemon, error) {
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

func (c *Client) CatchPokemon(pokemonName string, cache *pokecache.Cache) (SinglePokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint
	singlePokemon := SinglePokemon{}
	if data, ok := cache.GetCache(fullURL); ok {
		fmt.Println("cache hit")
		if err := json.Unmarshal(data, &singlePokemon); err != nil {
			return SinglePokemon{}, err
		}
		experience := singlePokemon.BaseExperience
		randInt := rand.IntN(experience)
		const threshold = 50
		if randInt < threshold {
			fmt.Printf("%v Caught\n", singlePokemon.Name)
			return singlePokemon, nil
		} else {
			fmt.Printf("%v not Caught\n", singlePokemon.Name)
			return SinglePokemon{}, nil
		}
	}
	res, err := c.httpClient.Get(fullURL)
	if err != nil {
		fmt.Println(err)
		return SinglePokemon{}, err
	}
	if res.StatusCode > 299 {
		fmt.Printf("Error status code: %v\n", res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return SinglePokemon{}, err
	}
	if err := json.Unmarshal(body, &singlePokemon); err != nil {
		fmt.Println(err)
		return SinglePokemon{}, err
	}

	experience := singlePokemon.BaseExperience
	randInt := rand.IntN(experience)
	const threshold = 50
	if randInt < threshold {
		fmt.Printf("%v Caught\n", singlePokemon.Name)

	} else {
		fmt.Printf("%v not Caught\n", singlePokemon.Name)
	}
	cache.AddCache(fullURL, body)
	return singlePokemon, nil
}

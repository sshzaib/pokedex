package main

import (
	"time"

	"github.com/sshzaib/pokedex/external/pokeapi"
	"github.com/sshzaib/pokedex/external/pokecache"
)

type config struct {
	cache               pokecache.Cache
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.SinglePokemon
}

func main() {
	interval := 1 * time.Hour
	config := &config{
		pokeapiClient: pokeapi.NewClient(),
		cache:         *pokecache.NewCache(interval),
		caughtPokemon: make(map[string]pokeapi.SinglePokemon),
	}
	StartRepl(config)
}

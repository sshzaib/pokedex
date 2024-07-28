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
}

func main() {
	interval := 3 * time.Second
	config := &config{
		pokeapiClient: pokeapi.NewClient(),
		cache:         *pokecache.NewCache(interval),
	}
	StartRepl(config)
}

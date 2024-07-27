package main

import (
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
	config := config{
		pokeapiClient: pokeapi.NewClient(),
		cache:         pokecache.NewCache(),
	}
	StartRepl(&config)
}

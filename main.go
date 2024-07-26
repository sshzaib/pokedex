package main

import "github.com/sshzaib/pokedex/external/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	config := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	StartRepl(&config)
}

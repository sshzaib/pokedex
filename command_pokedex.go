package main

import "fmt"

func commandPokedex(config *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex")
	if len(config.caughtPokemon) == 0 {
		fmt.Println("--no Pokemons caught--")
		return nil
	}
	for pokemonName := range config.caughtPokemon {
		fmt.Printf(" - %v\n", pokemonName)
	}
	return nil
}

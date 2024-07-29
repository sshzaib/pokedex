package main

import "fmt"

func commandInspect(config *config, args ...string) error {
	pokemonName := args[1]

	pokemon, ok := config.caughtPokemon[pokemonName]
	if !ok {
		fmt.Printf("you have not caught : %v\n", pokemonName)
		return nil
	}
	fmt.Print("\nPokemon INFO\n")
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Height)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Print("\nPokemon Stats\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %v:%v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Print("\nPokemon Types\n")
	for _, stat := range pokemon.Types {
		fmt.Printf(" - %v\n", stat.Type.Name)
	}
	return nil
}

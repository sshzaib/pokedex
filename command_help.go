package main

import "fmt"

func commandHelp(config *config, args ...string) error {
	commands := getCommands()
	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	for key, value := range commands {
		fmt.Printf("- %v: %v\n", key, value.description)
	}
	return nil
}

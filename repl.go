package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type cliCommands struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func StartRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >")
		scanner.Scan()
		text := scanner.Text()
		commandName := formatCommand(text)
		commands := getCommands()
		if len(commandName) == 2 {
			command, ok := commands[commandName[0]]
			if !ok {
				log.Fatal("invalid command: Enter help command to see all the valid commands")
			}
			command.callback(config, commandName...)
			continue
		}
		if len(commandName) == 0 {
			continue
		}
		command, ok := commands[commandName[0]]
		if !ok {
			log.Fatal("invalid command: Enter help command to see all the valid commands")
		}
		command.callback(config)
	}
}

func getCommands() map[string]cliCommands {
	cliCammands := map[string]cliCommands{
		"help": {
			name:        "help",
			description: "display the help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays Next Pokimon Locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display Previous Pokemon Locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display Pokemons in The Location",
			callback:    commandExplore,
		},
	}
	return cliCammands
}

func formatCommand(command string) []string {
	commandLower := strings.ToLower(command)
	commandSlice := strings.Split(commandLower, " ")
	return commandSlice
}

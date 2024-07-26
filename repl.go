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
	callback    func(*config) error
}

func StartRepl(config *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex >")
		scanner.Scan()
		commandName := scanner.Text()
		commandName = formatCommand(commandName)
		if len(commandName) == 0 {
			continue
		}
		commands := getCommands()
		command, ok := commands[commandName]
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
	}
	return cliCammands
}

func formatCommand(command string) string {
	commandLower := strings.ToLower(command)
	return commandLower
}

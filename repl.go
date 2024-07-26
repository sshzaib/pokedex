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
	callback    func() error
}

func StartRepl() {
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
		command.callback()
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
			description: "Displays 20 Pokimon Locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
	return cliCammands
}

func formatCommand(command string) string {
	commandLower := strings.ToLower(command)
	return commandLower
}

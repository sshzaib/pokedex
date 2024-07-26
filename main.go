package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type cliCommands struct {
	name        string
	description string
	callback    func() error || mapFunc
}

type result struct {
	name string
	url  string
}

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!\nUsage:\n\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

func commandExit() error {
	return nil
}

func commandMap(mapCount *int) error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatal("Can not print Pokimon locations")
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal(err)
	}
	for i := 0; i < 20; i++ {
		println(result.Results[i].Name)
	}
	return nil
}

type mapFunc func(*int)

func main() {
	cliCammand := map[string]cliCommands{
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
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("pokedex >")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error", err)
		}
		var mapInt int
		switch command {
		case "help\n":
			cliCammand["help"].callback()
		case "exit\n":
			cliCammand["exit"].callback()
			return
		case "map\n":
			cliCammand["map"].callback(&mapInt)
		}
	}
}

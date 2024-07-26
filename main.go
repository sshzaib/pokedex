package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Response struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap() error {
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

func main() {
	StartRepl()
}

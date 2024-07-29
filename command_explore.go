package main

func commandExplore(config *config, args ...string) error {
	client := config.pokeapiClient
	cache := &config.cache
	_, err := client.GetPokemonInLocation(args, cache)
	if err != nil {
		return err
	}
	return nil
}

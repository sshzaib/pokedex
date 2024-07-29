package main

func commandCatch(config *config, args ...string) error {
	cache := &config.cache
	pokemonName := args[1]
	res, err := config.pokeapiClient.CatchPokemon(pokemonName, cache)
	if err != nil {
		return err
	}
	config.caughtPokemon[pokemonName] = res
	return nil
}

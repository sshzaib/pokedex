package main

func commandMap(config *config, args ...string) error {
	client := config.pokeapiClient
	cache := &config.cache
	res, err := client.GetLocationAreas(config.nextLocationAreaURL, cache)
	if err != nil {
		return err
	}
	config.nextLocationAreaURL = res.Next
	config.prevLocationAreaURL = res.Previous
	return nil
}

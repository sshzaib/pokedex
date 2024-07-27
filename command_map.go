package main

func commandMap(config *config) error {
	client := config.pokeapiClient
	cache := config.cache
	res, err := client.GetLocationAreas(config.nextLocationAreaURL, cache)
	if err != nil {
		return err
	}
	config.nextLocationAreaURL = res.Next
	config.prevLocationAreaURL = res.Previous
	return nil
}

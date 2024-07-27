package main

func commandMapb(config *config) error {
	client := config.pokeapiClient
	cache := config.cache
	res, err := client.GetLocationAreas(config.prevLocationAreaURL, cache)
	if err != nil {
		return err
	}
	config.nextLocationAreaURL = res.Next
	config.prevLocationAreaURL = res.Previous
	return nil
}

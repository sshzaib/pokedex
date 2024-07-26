package main

func commandMap(config *config) error {
	client := config.pokeapiClient
	res, err := client.GetLocationAreas(config.nextLocationAreaURL)
	if err != nil {
		return err
	}
	config.nextLocationAreaURL = res.Next
	config.prevLocationAreaURL = res.Previous
	return nil
}

package main

func commandMapb(config *config) error {
	client := config.pokeapiClient
	res, err := client.GetLocationAreas(config.prevLocationAreaURL)
	if err != nil {
		return err
	}
	config.nextLocationAreaURL = res.Next
	config.prevLocationAreaURL = res.Previous
	return nil
}

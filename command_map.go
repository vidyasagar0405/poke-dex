package main

import (
	"fmt"

	"github.com/vidyasagar0405/pokedexcli/internals/pokeapi"
)

func callbackMap(cfg *config, args ...string) error {
	err := Map(cfg)
	if err != nil {
		return err
	}

	return nil
}

func Map(cfg *config) error {
	url := pokeapi.BaseURL + "/location-area"

	// Use next URL if available
	if cfg.nextLocationUrl != nil {
		url = *cfg.nextLocationUrl
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(url)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationUrl = resp.Next
	cfg.prevLocationUrl = resp.Prev

	return nil
}

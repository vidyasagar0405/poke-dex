package main

import (
	"fmt"

	"github.com/vidyasagar0405/pokedexcli/internals/pokeapi"
)

func callbackMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		return err
	}

	fmt.Println("locaiton Areas")

	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	return nil
}

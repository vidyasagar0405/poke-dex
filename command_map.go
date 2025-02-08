package main

import (
	"fmt"

	"github.com/vidyasagar0405/pokedexcli/internals/pokeapi"
)

func callbackMap() error {
    err := Map()
    if err!=nil {
		return err
    }
    
	return nil
}

func Map() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		return err
	}

	fmt.Println("locaiton Areas")

	for _, area := range resp.Results {
		fmt.Printf(" - %s, %s\n", area.Name, *resp.Next)
	}
    
	return nil
}

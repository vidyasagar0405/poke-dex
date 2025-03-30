package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}

	exporeResp, err := cfg.pokeapiClient.ExploreLocationArea(args[0])
	if err != nil {
		return err
	}

    fmt.Println("Found Pokemon:")
	for _, pokemon := range exporeResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)

	}

	return nil
}

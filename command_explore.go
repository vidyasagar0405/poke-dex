package main

import "fmt"

func callbackExplore(cfg *config, locArea string) error {

	exporeResp, err := cfg.pokeapiClient.ExploreLocationArea(locArea)
	if err != nil {
		return err
	}

    fmt.Println("Found Pokemon:")
	for _, pokemon := range exporeResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)

	}

	return nil
}

package cmds

import (
	"fmt"

	"github.com/vidyasagar0405/pokedexcli/config"
	"github.com/vidyasagar0405/pokedexcli/internals/pokeapi"
)

func callbackMap(cfg *config.Config, args ...string) error {
	err := Map(cfg)
	if err != nil {
		return err
	}

	return nil
}

func Map(cfg *config.Config) error {
	url := pokeapi.BaseURL + "/location-area"

	// Use next URL if available
	if cfg.NextLocationUrl != nil {
		url = *cfg.NextLocationUrl
	}

	resp, err := cfg.PokeapiClient.ListLocationAreas(url)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.NextLocationUrl = resp.Next
	cfg.PrevLocationUrl = resp.Prev

	return nil
}

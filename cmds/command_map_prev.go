package cmds

import (
	"fmt"

	"github.com/vidyasagar0405/pokedexcli/config"
)

func callbackMapPrev(cfg *config.Config, args ...string) error {
	err := MapPrev(cfg)
	if err != nil {
		return err
	}

	return nil
}

func MapPrev(cfg *config.Config) error {
	if cfg.PrevLocationUrl == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	url := *cfg.PrevLocationUrl

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

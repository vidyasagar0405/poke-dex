package main

import (
	"fmt"
)

func callbackMapPrev(cfg *config) error {
	err := MapPrev(cfg)
	if err != nil {
		return err
	}

	return nil
}

func MapPrev(cfg *config) error {
	if cfg.prevLocationUrl == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	url := *cfg.prevLocationUrl

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

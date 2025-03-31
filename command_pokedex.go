package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {

    if len(cfg.caughtPokemon) == 0 {
        fmt.Println("You have not caught any pokemon")
        return nil
    }

    fmt.Println("Your Pokedex:")

	for _, val := range cfg.caughtPokemon {
		fmt.Printf(" - %s", val.Name)
	}
    return nil

}

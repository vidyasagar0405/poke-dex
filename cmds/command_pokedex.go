package cmds

import (
	"fmt"

	"github.com/vidyasagar0405/pokedexcli/config"
)

func callbackPokedex(cfg *config.Config, args ...string) error {

    if len(cfg.CaughtPokemon) == 0 {
        fmt.Println("You have not caught any pokemon")
        return nil
    }

    fmt.Println("Your Pokedex:")

	for _, val := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", val.Name)
	}
    return nil

}

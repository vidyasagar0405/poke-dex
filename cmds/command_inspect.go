package cmds

import (
	"errors"
	"fmt"

	"github.com/vidyasagar0405/pokedexcli/config"
)

func callbackInspect(cfg *config.Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	pokemon, exists := cfg.CaughtPokemon[name]
	if !exists {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, val := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, val := range pokemon.Types {
		fmt.Printf(" - %s\n", val.Type.Name)
	}

	return nil
}

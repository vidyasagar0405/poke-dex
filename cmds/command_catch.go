package cmds

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/vidyasagar0405/pokedexcli/config"
)

func callbackCatch(cfg *config.Config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.PokeapiClient.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
    fmt.Println("You may now inspect it with the inspect command.")

	cfg.CaughtPokemon[pokemon.Name] = pokemon
	return nil
}

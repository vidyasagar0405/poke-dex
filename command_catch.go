package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(name)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

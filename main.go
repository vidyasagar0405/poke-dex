package main

import (
	"time"

	"github.com/vidyasagar0405/pokedexcli/internals/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	caughtPokemon   map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}

	startRepl(&cfg)
}

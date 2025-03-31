package main

import (
	"time"

	"github.com/vidyasagar0405/pokedexcli/internals/pokeapi"
	"github.com/vidyasagar0405/pokedexcli/config"
)

func main() {

	cfg := &config.Config{
		CaughtPokemon: map[string]pokeapi.Pokemon{},
		PokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}

	startRepl(cfg)
}

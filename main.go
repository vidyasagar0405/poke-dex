package main

import "github.com/vidyasagar0405/pokedexcli/internals/pokeapi"

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}

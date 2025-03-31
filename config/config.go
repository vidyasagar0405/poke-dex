package config

import "github.com/vidyasagar0405/pokedexcli/internals/pokeapi"

type Config struct {
	PokeapiClient   pokeapi.Client
	NextLocationUrl *string
	PrevLocationUrl *string
	CaughtPokemon   map[string]pokeapi.Pokemon
}

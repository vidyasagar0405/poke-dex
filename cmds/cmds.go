package cmds

import "github.com/vidyasagar0405/pokedexcli/config"

type clicommand struct {
	Callback    func(cfg *config.Config, arg ...string) error
	Name        string
	Description string
}

func GetCommands() map[string]clicommand {
	return map[string]clicommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    callbackHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    callbackExit,
		},
		"map": {
			Name:        "map",
			Description: "lists available location area, type again to see the next page",
			Callback:    callbackMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "lists previous page of location area",
			Callback:    callbackMapPrev,
		},
		"explore": {
			Name:        "explore",
			Description: "list all the pokemons in a area",
			Callback:    callbackExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "catch a pokemon",
			Callback:    callbackCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "inspect a caught pokemon",
			Callback:    callbackInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "lists all caught pokemon",
			Callback:    callbackPokedex,
		},
	}
}

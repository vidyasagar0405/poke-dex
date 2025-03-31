package main

type clicommand struct {
	callback    func(cfg *config, arg ...string) error
	name        string
	description string
}

func getCommands() map[string]clicommand {
	return map[string]clicommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"map": {
			name:        "map",
			description: "lists available location area, type again to see the next page",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "lists previous page of location area",
			callback:    callbackMapPrev,
		},
		"explore": {
			name:        "explore",
			description: "list all the pokemons in a area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch",
			description: "catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "inspect a caught pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "lists all caught pokemon",
			callback:    callbackPokedex,
		},
	}
}

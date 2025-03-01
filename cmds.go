package main

type clicommand struct {
	callback    func(cfg *config) error
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
		"map-prev": {
			name:        "map-prev",
			description: "lists previous page of location area",
			callback:    callbackMapPrev,
		},
	}
}

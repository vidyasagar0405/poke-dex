package cmds

import (
	"fmt"

	"github.com/vidyasagar0405/pokedexcli/config"
)

func callbackHelp(cfg *config.Config, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	availableCommand := GetCommands()
	for _, cmd := range availableCommand {
		fmt.Printf(" - %s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()
	return nil
}

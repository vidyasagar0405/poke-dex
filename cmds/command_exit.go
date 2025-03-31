package cmds

import (
	"os"

	"github.com/vidyasagar0405/pokedexcli/config"
)

func callbackExit(cfg *config.Config, args ...string) error {
	os.Exit(0)
	return nil
}

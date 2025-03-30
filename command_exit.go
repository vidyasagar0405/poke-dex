package main

import "os"

func callbackExit(cfg *config, _ string) error {
	os.Exit(0)
	return nil
}

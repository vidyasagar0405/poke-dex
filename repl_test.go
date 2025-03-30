package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Run tests and return exit code
	code := m.Run()
	os.Exit(code)
}

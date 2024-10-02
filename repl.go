package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokedex > ")

		scanner.Scan()
		raw_input_text := scanner.Text()

		cleaned_input := cleanInput(raw_input_text)

		if len(cleaned_input) == 0 {
			continue
		}

		commandName := cleaned_input[0]

		available_command := getCommands()
		command, ok := available_command[commandName]

		if !ok {
			fmt.Println("Invalid command")
		}

		command.callback()

	}
}

func cleanInput(raw_input_text string) []string {
	lowered := strings.ToLower(raw_input_text)
	words := strings.Fields(lowered)
	return words
}

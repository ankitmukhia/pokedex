package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ankitmukhia/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next *string
	prev *string
}

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if reader.Scan() {
			word := reader.Text()
			input := cleanInput(word)

			firstWord := input[0]
			command, exists := command()[firstWord]

			if exists {
				err := command.callback()
				// this loop shouldn't stop even command are not valid, so continue will make sure that. To continue running it
				if err != nil {
					fmt.Print(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}

func cleanInput(test string) []string {
	trimed := strings.ToLower(test)
	output := strings.Fields(trimed)
	return output
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func command() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Help the Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Log pakedex location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Log pakedex previous location",
			callback:    commandMapB,
		},
	}
}

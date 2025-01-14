package main

import (
	"fmt"
)

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range command() {
		fmt.Printf("%s\n: %s\n", c.name, c.description)
	}
	return nil
}

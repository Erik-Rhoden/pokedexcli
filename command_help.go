package main

import "fmt"

func commandHelp() error {
	fmt.Printf(`
Welcome to the Pokedex!
Usage:

`)
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

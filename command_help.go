package main

import "fmt"

func commandHelp(config *Config) error {
	commands := getCliCommands()

	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, command := range commands {
		fmt.Printf("%s:\t %s\n", command.name, command.description)
	}
	fmt.Println("")
	return nil
}

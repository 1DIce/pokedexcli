package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
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

func commandExit() error {
	os.Exit(0)
	return nil
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    func() error { return commandHelp() },
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    func() error { return commandExit() },
		},
	}
}

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)

	availableCommands := getCliCommands()

	for {
		fmt.Print("Pokedex > ")
		hasMoreTokens := inputScanner.Scan()
		if hasMoreTokens {
			// TODO handle this later?!
		}

		input := inputScanner.Text()

		command, ok := availableCommands[input]
		if !ok {
			fmt.Printf("'%s' is not a valid command!\n", input)
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Printf("Error during command execution %v\n", err)
		}

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapB,
		},
	}
}

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)

	availableCommands := getCliCommands()
	config := Config{
		LocationArea: LocationAreaConfig{},
	}

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

		err := command.callback(&config)
		if err != nil {
			fmt.Printf("Error during command execution %v\n", err)
		}

	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, arguments []string) error
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
		"explore": {
			name:        "explore",
			description: "Lists all Pokemon in a given area or area id",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch Pokemon with a given name or id. Caught Pokemon are added to the Pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught Pokemon in the Pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all Pokemon in the Pokedex",
			callback:    commandPokedex,
		},
	}
}

func main() {
	inputScanner := bufio.NewScanner(os.Stdin)

	availableCommands := getCliCommands()
	config := NewConfig()

	for {
		fmt.Print("Pokedex > ")
		hasMoreTokens := inputScanner.Scan()
		if hasMoreTokens {
			// TODO handle this later?!
		}

		input := inputScanner.Text()

		arguments := strings.Split(input, " ")
		command, ok := availableCommands[arguments[0]]
		if !ok {
			fmt.Printf("'%s' is not a valid command!\n", input)
			continue
		}

		err := command.callback(&config, arguments[1:])
		if err != nil {
			fmt.Printf("Error during command execution %v\n", err)
		}

	}
}

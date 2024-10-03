package main

import (
	"fmt"
	"log"

	pokeapi "github.com/1DIce/pokedexcli/poke_api"
)

func commandExplore(config *Config, arguments []string) error {
	if config == nil {
		log.Fatal("config was nil")
	}

	fmt.Println("")

	if len(arguments) == 0 {
		fmt.Println("Error: No location name given")
	}

	if len(arguments) > 1 {
		fmt.Println("Error: To many location names given! Input only a single location name")
	}

	locationNameOrId := arguments[0]
	response, err := pokeapi.FetchLocation(locationNameOrId)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", response.Name)

	for _, encounter := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	fmt.Println("")
	return nil
}

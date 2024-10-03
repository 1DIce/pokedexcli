package main

import (
	"fmt"
	"log"
	"math/rand"

	pokeapi "github.com/1DIce/pokedexcli/poke_api"
)

func commandCatch(config *Config, arguments []string) error {
	if config == nil {
		log.Fatal("config was nil")
	}

	if len(arguments) == 0 {
		fmt.Println("Error: No location name given")
	}

	if len(arguments) > 1 {
		fmt.Println("Error: To many location names given! Input only a single location name")
	}

	pokemonNameOrId := arguments[0]
	response, err := pokeapi.FetchPokemonDetails(pokemonNameOrId)
	if err != nil {
		return err
	}

	isCatchSuccess := rand.Int63n(int64(response.BaseExperience)) < 80

	if isCatchSuccess {
		fmt.Printf("Successfully caught %s!!!\n\n", response.Name)
		config.Pokedex[response.Name] = response
	} else {
		fmt.Printf("Failed to catch %s... Try again!\n\n", response.Name)
	}

	return nil
}

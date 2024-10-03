package main

import (
	"fmt"
	"log"

	pokeapi "github.com/1DIce/pokedexcli/poke_api"
)

func commandInspect(config *Config, arguments []string) error {
	if config == nil {
		log.Fatal("config was nil")
	}

	if len(arguments) == 0 {
		fmt.Println("Error: No pokemon name")
	}

	if len(arguments) > 1 {
		fmt.Println("Error: To many arguments given! Input only a single pokemon name")
	}

	pokemonName := arguments[0]
	details, exists := config.Pokedex[pokemonName]
	if !exists {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}

	printPokemonInfos(details)
	fmt.Println("")

	return nil
}

func printPokemonInfos(details pokeapi.PokemonDetails) {
	fmt.Printf(`
Name: %s
Height: %d
Weight: %d
`, details.Name, details.Height, details.Weight)

	fmt.Println("Stats:")
	for _, stat := range details.Stats {
		fmt.Printf(" - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range details.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
}

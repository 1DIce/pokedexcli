package main

import (
	"fmt"
	"log"
)

func commandPokedex(config *Config, arguments []string) error {
	if config == nil {
		log.Fatal("config was nil")
	}

	fmt.Printf("\n%d Pokemon in Pokedex:\n", len(config.Pokedex))

	for _, pokemon := range config.Pokedex {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	fmt.Println("")

	return nil
}

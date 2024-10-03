package main

import (
	"fmt"
	"log"

	pokeapi "github.com/1DIce/pokedexcli/poke_api"
)

func commandMap(config *Config) error {
	if config == nil {
		log.Fatal("config was nil")
	}

	currentPage := config.LocationArea.CurrentPageIndex
	nextPage := uint(0)
	if currentPage != nil {
		nextPage = *currentPage + 1
	}

	if nextPage > config.LocationArea.LastPageIndex {
		fmt.Println("No more pages available")
		fmt.Println("")
	}

	response, err := pokeapi.GetLocationAreas(nextPage)
	if err != nil {
		return err
	}

	config.LocationArea.CurrentPageIndex = &response.CurrentPageIndex
	config.LocationArea.LastPageIndex = response.LastPageIndex

	fmt.Printf("\nLocations page %d of %d:\n\n", response.CurrentPageIndex+1, response.LastPageIndex+1)

	for _, result := range response.Areas {
		fmt.Println(result.Name)
	}

	fmt.Println("")
	return nil
}

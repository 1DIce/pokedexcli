package main

import (
	"fmt"
	"log"

	pokeapi "github.com/1DIce/pokedexcli/poke_api"
)

func commandMap(config *Config, arguments []string) error {
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

	response, err := pokeapi.FetchLocationAreas(nextPage)
	if err != nil {
		return err
	}

	config.LocationArea.CurrentPageIndex = &response.CurrentPageIndex
	config.LocationArea.LastPageIndex = response.LastPageIndex

	fmt.Printf("\nLocations page %d of %d:\n\n", response.CurrentPageIndex+1, response.LastPageIndex+1)

	for _, result := range response.Areas {
		fmt.Printf("%d - %s\n", result.Id, result.Name)
	}

	fmt.Println("")
	return nil
}

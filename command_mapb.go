package main

import (
	"fmt"
	"log"

	pokeapi "github.com/1DIce/pokedexcli/poke_api"
)

func commandMapB(config *Config, arguments []string) error {
	if config == nil {
		log.Fatal("config was nil")
	}

	currentPage := config.LocationArea.CurrentPageIndex
	if currentPage == nil || *currentPage == 0 {
		fmt.Println("No previous page available")
		fmt.Println("")
		return nil
	}
	previousPage := *currentPage - 1

	response, err := pokeapi.FetchLocationAreas(previousPage)
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

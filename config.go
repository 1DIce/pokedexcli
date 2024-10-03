package main

import (
	pokeapi "github.com/1DIce/pokedexcli/poke_api"
)

type LocationAreaConfig struct {
	CurrentPageIndex *uint
	LastPageIndex    uint
}

type Config struct {
	LocationArea LocationAreaConfig
	Pokedex      map[string]pokeapi.PokemonDetails
}

func NewConfig() Config {
	config := Config{
		LocationArea: LocationAreaConfig{},
		Pokedex:      make(map[string]pokeapi.PokemonDetails),
	}
	return config
}

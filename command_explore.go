package main

import (
	"fmt"
	"time"

	"github.com/csullivan94/pokedex/internal/pokeapi"
)

func commandExplore(cfg *Config) error {
	start := time.Now()
	if cfg.Argument == "" {
		fmt.Println("Missing exploration location")
		return fmt.Errorf("no argument given")
	}

	url := "https://pokeapi.co/api/v2/location-area/" + cfg.Argument

	fmt.Println(url)

	locationAreaStruct, err := pokeapi.GetLocationPokemon(url, cfg.Cache)
	if err != nil {
		return err
	}

	for item := range locationAreaStruct.PokemonEncounters {
		fmt.Println(locationAreaStruct.PokemonEncounters[item].Pokemon.Name)

	}
	if pokeapi.CacheUsed {
		fmt.Println("Cache used: ", time.Since(start).Seconds())
	} else {
		fmt.Println("Cache not used: ", time.Since(start).Seconds())
	}
	pokeapi.CacheUsed = false

	return nil
}

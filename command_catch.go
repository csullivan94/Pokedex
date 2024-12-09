package main

import (
	"fmt"

	"github.com/csullivan94/pokedex/internal/pokeapi"
)

func commandCatch(cfg *Config) error {
	if cfg.Argument == "" {
		fmt.Println("No pokemon to catch!")
		return fmt.Errorf("no pokemon argument specified")
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + cfg.Argument

	pokemonStruct, err := pokeapi.GetPokemonDetails(url, cfg.Cache)
	if err != nil {
		return err
	}

	fmt.Printf("You caught %s!\n", cfg.Argument)
	fmt.Printf("Adding %s to pokedex...\n", cfg.Argument)
	fmt.Println("Name: ", pokemonStruct.Name)
	fmt.Println("Height: ", pokemonStruct.Height)
	fmt.Println("Weight: ", pokemonStruct.Weight)

	cfg.Pokedex[cfg.Argument] = pokemonStruct

	return nil
}

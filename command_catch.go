package main

import (
	"fmt"
	"math/rand/v2"

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

	fmt.Printf("Throwing a Pokeball at %s...", cfg.Argument)

	if !randCatch(pokemonStruct) {
		fmt.Printf("%s escaped!\n", pokemonStruct.Name)
	} else {

		fmt.Printf("You caught %s!\n", cfg.Argument)

		cfg.Pokedex[cfg.Argument] = pokemonStruct
	}

	return nil
}

func randCatch(pokemonStruct pokeapi.Pokemon) bool {
	baseExperience := pokemonStruct.BaseExperience / 2
	if baseExperience > 250 {
		baseExperience = 250
	}
	chance := (308.0 - float32(baseExperience)) / 308.0
	return rand.Float32() < chance
}

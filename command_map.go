package main

import (
	"fmt"

	"github.com/csullivan94/pokedex/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location"
	}
	cfg.PageNum += 1

	locationstruct, err := pokeapi.GetLocations(cfg.Next)
	if err != nil {
		return err
	}

	for item := range locationstruct.Results {
		fmt.Println(locationstruct.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", cfg.PageNum)
	cfg.Next = locationstruct.Next
	if cfg.Next == "" {
		fmt.Println("End of locations")
		cfg.PageNum = 0
	}
	cfg.Previous = locationstruct.Previous

	return nil
}

func commandMapb(cfg *Config) error {

	if cfg.Previous == "" {
		fmt.Println("No previous pages")
		return fmt.Errorf("no previous pages")
	}

	cfg.PageNum -= 1

	locationstruct, err := pokeapi.GetLocations(cfg.Previous)
	if err != nil {
		return err
	}

	for item := range locationstruct.Results {
		fmt.Println(locationstruct.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", cfg.PageNum)
	cfg.Next = locationstruct.Next
	cfg.Previous = locationstruct.Previous

	return nil
}

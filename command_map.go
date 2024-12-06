package main

import (
	"fmt"
	"time"

	"github.com/csullivan94/pokedex/internal/pokeapi"
	"github.com/csullivan94/pokedex/internal/pokecache"
)

func commandMap(cfg *Config, cache *pokecache.Cache) error {
	start := time.Now()

	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location?offset=0&limit=20"
	}
	cfg.PageNum += 1

	cfg.Current = cfg.Next

	locationstruct, err := pokeapi.GetLocations(cfg.Current, cache)
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

	if pokeapi.CacheUsed {
		fmt.Println("Cache used: ", time.Since(start).Seconds())
	} else {
		fmt.Println("Cache not used: ", time.Since(start).Seconds())
	}
	pokeapi.CacheUsed = false

	return nil
}

func commandMapb(cfg *Config, cache *pokecache.Cache) error {

	start := time.Now()

	if cfg.Previous == "" {
		fmt.Println("No previous pages")
		return fmt.Errorf("no previous pages")
	}

	cfg.PageNum -= 1
	cfg.Current = cfg.Previous

	locationstruct, err := pokeapi.GetLocations(cfg.Current, cache)
	if err != nil {
		return err
	}

	for item := range locationstruct.Results {
		fmt.Println(locationstruct.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", cfg.PageNum)
	cfg.Next = locationstruct.Next
	cfg.Previous = locationstruct.Previous

	if pokeapi.CacheUsed {
		fmt.Println("Cache used: ", time.Since(start).Seconds())
	} else {
		fmt.Println("Cache not used: ", time.Since(start).Seconds())
	}
	pokeapi.CacheUsed = false
	return nil
}

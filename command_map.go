package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/csullivan94/pokedex/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	start := time.Now()

	if cfg.CurrentMap == "" {
		cfg.CurrentMap = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	} else {
		if cfg.NextMap != "" {
			cfg.CurrentMap = cfg.NextMap
		} else {
			cfg.CurrentMap = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
		}

	}

	if cfg.Argument != "" {
		pageNum, err := strconv.Atoi(cfg.Argument)
		if err != nil {
			return err
		}
		err = GivePageNumber(cfg, pageNum)
		if err != nil {
			return err
		}

	}
	GetPageNumber(cfg)

	fmt.Println(cfg.CurrentMap)

	locationstruct, err := pokeapi.GetLocations(cfg.CurrentMap, cfg.Cache)
	if err != nil {
		return err
	}

	for item := range locationstruct.Results {
		fmt.Println(locationstruct.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", cfg.PageNum)
	cfg.NextMap = locationstruct.Next
	if cfg.NextMap == "" {
		fmt.Println("End of locations")
		cfg.PageNum = 0
	}
	cfg.PreviousMap = locationstruct.Previous

	if pokeapi.CacheUsed {
		fmt.Println("Cache used: ", time.Since(start).Seconds())
	} else {
		fmt.Println("Cache not used: ", time.Since(start).Seconds())
	}
	pokeapi.CacheUsed = false

	return nil
}

func commandMapb(cfg *Config) error {

	start := time.Now()

	if cfg.PreviousMap == "" {
		fmt.Println("No previous pages")
		return fmt.Errorf("no previous pages")
	}

	cfg.CurrentMap = cfg.PreviousMap
	GetPageNumber(cfg)

	locationstruct, err := pokeapi.GetLocations(cfg.CurrentMap, cfg.Cache)
	if err != nil {
		return err
	}

	for item := range locationstruct.Results {
		fmt.Println(locationstruct.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", cfg.PageNum)
	cfg.NextMap = locationstruct.Next
	cfg.PreviousMap = locationstruct.Previous

	if pokeapi.CacheUsed {
		fmt.Println("Cache used: ", time.Since(start).Seconds())
	} else {
		fmt.Println("Cache not used: ", time.Since(start).Seconds())
	}
	pokeapi.CacheUsed = false
	return nil
}

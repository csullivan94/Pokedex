package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/csullivan94/pokedex/internal/pokeapi"
)

func commandMap(cfg *Config) error {
	start := time.Now()

	if cfg.Current == "" {
		cfg.Current = "https://pokeapi.co/api/v2/location?offset=0&limit=20"
	} else {
		if cfg.Next != "" {
			cfg.Current = cfg.Next
		} else {
			cfg.Current = "https://pokeapi.co/api/v2/location?offset=0&limit=20"
		}

	}

	if cfg.Argument != "" {
		pageNum, err := strconv.Atoi(cfg.Argument)
		if err != nil {
			return err
		}
		err = page_number.GivePageNumber(cfg, pageNum)
		if err != nil {
			return err
		}

	}
	getPageNumber(cfg)

	fmt.Println(cfg.Current)

	locationstruct, err := pokeapi.GetLocations(cfg.Current, cfg.Cache)
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
	cfg.Argument = ""
	return nil
}

func commandMapb(cfg *Config) error {

	start := time.Now()

	if cfg.Previous == "" {
		fmt.Println("No previous pages")
		return fmt.Errorf("no previous pages")
	}

	cfg.Current = cfg.Previous
	getPageNumber(cfg)

	locationstruct, err := pokeapi.GetLocations(cfg.Current, cfg.Cache)
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

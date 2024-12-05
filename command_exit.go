package main

import (
	"fmt"
	"os"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

func commandExit(cfg *Config, cache *pokecache.Cache) error {
	fmt.Println("exiting the Pokedex...")
	os.Exit(0)
	return nil
}

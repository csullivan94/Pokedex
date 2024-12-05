package main

import (
	"fmt"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

func commandHelp(cfg *Config, cache *pokecache.Cache) error {
	fmt.Println("----------Welcome to the Pokedex----------")
	for _, commandStruct := range getCommands() {
		fmt.Println("Command: ", commandStruct.name)
		fmt.Println(commandStruct.description)
		fmt.Println("------------------------------")
	}
	return nil
}

package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config) error {
	fmt.Println("exiting the Pokedex...")
	os.Exit(0)
	return nil
}

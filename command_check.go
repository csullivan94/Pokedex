package main

import "fmt"

func commandCheck(cfg *Config) error {
	for pokemon := range cfg.Pokedex {
		fmt.Println(cfg.Pokedex[pokemon].Name)
	}
	return nil
}

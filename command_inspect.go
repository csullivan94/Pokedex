package main

import "fmt"

func commandInspect(cfg *Config) error {
	if cfg.Argument == "" {
		fmt.Println("Your Pokedex:")
		for pokemon := range cfg.Pokedex {
			fmt.Println(cfg.Pokedex[pokemon].Name)
		}
		return nil
	} else {
		pokemon, exists := cfg.Pokedex[cfg.Argument]
		if !exists {
			fmt.Printf("No %s in your Pokedex\n", cfg.Argument)
			return nil
		}

		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", cfg.Pokedex[cfg.Argument].Height)
		fmt.Println("Weight:", cfg.Pokedex[cfg.Argument].Weight)
		fmt.Println("Stats:")
		for item := range cfg.Pokedex[cfg.Argument].Stats {
			fmt.Printf("	-%s: %v\n", cfg.Pokedex[cfg.Argument].Stats[item].Stat.Name, cfg.Pokedex[cfg.Argument].Stats[item].BaseStat)
		}
		fmt.Println("Types:")
		for item := range cfg.Pokedex[cfg.Argument].Types {
			fmt.Printf("	- %s\n", cfg.Pokedex[cfg.Argument].Types[item].Type.Name)
		}
		return nil
	}

}

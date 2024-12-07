package main

import "fmt"

func commandExplore(cfg *Config) error {
	if cfg.Argument == "" {
		fmt.Println("Missing exploration location")
		return fmt.Errorf("no argument given")
	}
	return nil
}

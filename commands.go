package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Short for 'map back'\nLists the previous 20 locations",
			callback:    commandMapb,
		},
	}

}

func commandHelp(cfg *Config) error {
	fmt.Println("----------Welcome to the Pokedex----------")
	for _, commandStruct := range getCommands() {
		fmt.Println("Command: ", commandStruct.name)
		fmt.Println(commandStruct.description)
		fmt.Println("------------------------------")
	}
	return nil
}

func commandExit(cfg *Config) error {
	fmt.Println("exiting the Pokedex...")
	os.Exit(0)
	return nil
}

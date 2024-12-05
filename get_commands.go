package main

import "github.com/csullivan94/pokedex/internal/pokecache"

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, *pokecache.Cache) error
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

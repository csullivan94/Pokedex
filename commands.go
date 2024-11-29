package main

import (
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": cliCommand{
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": cliCommand{
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("----------Welcome to the Pokedex----------")
	for _, commandStruct := range commandsMap() {
		fmt.Println("Command: ", commandStruct.name)
		fmt.Println(commandStruct.description)
		fmt.Println("------------------------------")
	}
	return nil
}

func commandExit() error {
	return nil
}

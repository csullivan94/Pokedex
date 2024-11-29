package commands

import (
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands map[string]cliCommand

func commandHelp() error {
	for _, commandStruct := range commands {
		fmt.Println(commandStruct.name)
		fmt.Println(commandStruct.description)
	}
	return nil
}

func commandExit() error {
	return nil
}

func initCommands() {
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exiting the Pokedex",
		callback:    commandExit,
	}
}

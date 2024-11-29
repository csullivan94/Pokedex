package main

import "fmt"

func commandHelp(cfg *Config) error {
	fmt.Println("----------Welcome to the Pokedex----------")
	for _, commandStruct := range getCommands() {
		fmt.Println("Command: ", commandStruct.name)
		fmt.Println(commandStruct.description)
		fmt.Println("------------------------------")
	}
	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
		"map": cliCommand{
			name:        "map",
			description: "Lists the next 20 locations",
			callback:    commandMap,
		},
		"mapb": cliCommand{
			name:        "mapb",
			description: "Short for 'map back'\nLists the previous 20 locations",
			callback:    commandMapb,
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
	fmt.Println("exiting the Pokedex...")
	os.Exit(0)
	return nil
}

var nextUrl string
var prevUrl string
var pageNum int

func commandMap() error {
	if nextUrl == "" {
		nextUrl = "https://pokeapi.co/api/v2/location"
	}

	pageNumPtr := &pageNum
	*pageNumPtr += 1

	res, err := http.Get(nextUrl)
	if err != nil {
		fmt.Println("error with get request")
		return err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locations Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Printf("error unmarshaling data: %v ", err)
		return err
	}

	for item := range locations.Results {
		fmt.Println(locations.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", pageNum)
	nextUrl = locations.Next
	if nextUrl == "" {
		fmt.Println("End of locations")
		*pageNumPtr = 0
	}
	prevUrl = locations.Previous

	return nil
}

func commandMapb() error {

	if prevUrl == "" {
		fmt.Println("No previous pages")
		return fmt.Errorf("no previous pages")
	}

	pageNumPtr := &pageNum
	*pageNumPtr -= 1

	res, err := http.Get(prevUrl)
	if err != nil {
		fmt.Println("error with get request")
		return err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locations Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Printf("error unmarshaling data: %v ", err)
		return err
	}

	for item := range locations.Results {
		fmt.Println(locations.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", pageNum)
	nextUrl = locations.Next
	prevUrl = locations.Previous

	return nil
}

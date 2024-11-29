package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *Config) error {

	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location"
	}

	cfg.PageNum += 1

	res, err := http.Get(cfg.Next)
	if err != nil {
		fmt.Println("error with get request: ", err)
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var locations Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Printf("error unmarshaling data: %v ", err)
		return err
	}

	for item := range locations.Results {
		fmt.Println(locations.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", cfg.PageNum)
	cfg.Next = locations.Next
	if cfg.Next == "" {
		fmt.Println("End of locations")
		cfg.PageNum = 0
	}
	cfg.Previous = locations.Previous

	return nil
}

func commandMapb(cfg *Config) error {

	if cfg.Previous == "" {
		fmt.Println("No previous pages")
		return fmt.Errorf("no previous pages")
	}

	cfg.PageNum -= 1

	res, err := http.Get(cfg.Previous)
	if err != nil {
		fmt.Println("error with get request")
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var locations Location
	err = json.Unmarshal(data, &locations)
	if err != nil {
		fmt.Printf("error unmarshaling data: %v ", err)
		return err
	}

	for item := range locations.Results {
		fmt.Println(locations.Results[item].Name)

	}
	fmt.Printf("---------------Page %d---------------\n", cfg.PageNum)
	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/csullivan94/pokedex/internal/pokeapi"
	"github.com/csullivan94/pokedex/internal/pokecache"
)

func main() {

	cfg := &Config{}
	interval := 10 * time.Second
	cfg.Cache = pokecache.NewCache(interval)
	cfg.Pokedex = make(map[string]pokeapi.Pokemon)

	scanner := bufio.NewScanner(os.Stdin)
	var input string

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input = scanner.Text()
		inputSplit := strings.Split(input, " ")

		if len(inputSplit) > 1 {
			cfg.Argument = inputSplit[1]
		}
		command, exists := getCommands()[inputSplit[0]]
		if !exists {
			fmt.Println("unknown command")
		} else {
			command.callback(cfg)
		}
		cfg.Argument = ""

	}

}

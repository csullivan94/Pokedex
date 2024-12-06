package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

func main() {

	cfg := &Config{}
	cache := pokecache.NewCache(5)

	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input = scanner.Text()

		command, exists := getCommands()[input]
		if !exists {
			fmt.Println("unknown command")
		} else {
			command.callback(cfg, cache)
		}

	}

}

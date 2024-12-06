package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/csullivan94/pokedex/internal/pokecache"
)

func main() {

	cfg := &Config{}
	interval := 10 * time.Second
	cache := pokecache.NewCache(interval)

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

		for value := range cache.Data {
			elapsed := time.Since(cache.Data[value].CreatedAt)
			fmt.Printf("%v time elapsed: %v\n", value, elapsed)

		}
	}

}

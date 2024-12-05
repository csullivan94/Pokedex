package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	cfg := &Config{}
	NewCache()

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
			command.callback(cfg)
		}

	}

}

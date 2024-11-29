package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input = scanner.Text()

		switch input {
		case "help":
			commandHelp()
		case "exit":
			fmt.Println("exiting the Pokedex...")
			os.Exit(0)
		default:
			fmt.Println("unknown command")
		}

	}
}

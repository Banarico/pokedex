package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
    "github.com/Banarico/pokedex/internal/pokeapi"
)

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		// Wait for the user to type something and press enter
		scanner.Scan()

		// Get the text the user typed
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		// Look up the command in your map of commands
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second)
    cfg := &Config{
        pokeapiClient: pokeClient,
    }
    startRepl(cfg)
}

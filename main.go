package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
    "github.com/Banarico/pokedex/internal/pokeapi"
    "github.com/Banarico/pokedex/internal/pokecache"
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
                arguments := words[1:]

		// Look up the command in your map of commands
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, arguments...)
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
    chc := pokecache.NewCache(5 * time.Second)
    pokeClient := pokeapi.NewClient(5 * time.Second, chc)
    cfg := &Config{
        pokeapiClient: pokeClient,
        cache:  chc,
    }
    startRepl(cfg)
}

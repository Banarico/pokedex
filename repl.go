package main

import (
    "strings"
    "os"
    "fmt"
    "github.com/Banarico/pokedex/internal/pokeapi"
    "github.com/Banarico/pokedex/internal/pokecache"
)

func cleanInput(text string) []string {
    lower := strings.ToLower(text)
    split := strings.Fields(lower)
    return split
}

func getCommands() map[string]cliCommand {
    cmds := map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "map": {
            name:        "map",
            description: "Displays next page of areas",
            callback:    commandMap,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays previous page of areas",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore",
            description: "A list of Pokemon for a specific area",
            callback:    commandExpl,
        },
    }
    return cmds
}

func commandExit(cfg *Config, arg ...string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp(cfg *Config, arg ...string) error {
    list := getCommands()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    for _, cmd := range list {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

type Config struct {
    pokeapiClient    pokeapi.Client
    cache            pokecache.Cache
    nextLocationsURL *string // use a pointer because these can be null
    prevLocationsURL *string
}

type List struct {
    Next     *string `json:"next"`
    Previous *string `json:"previous"`
    Results  []struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"results"`
}

func commandMap(cfg *Config, arg ...string) error {
    command, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
    if err != nil {
        fmt.Println(err)
    }
    cfg.nextLocationsURL = command.Next
    cfg.prevLocationsURL = command.Previous
    for _, a := range command.Results {
        fmt.Println(a.Name)
    }
    return nil
}

func commandMapb(cfg *Config, arg ...string) error {
    if cfg.prevLocationsURL == nil {
        fmt.Println("you're on the first page")
        return nil
    }
    command, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
    if err != nil {
        fmt.Println(err)
    }
    cfg.nextLocationsURL = command.Next
    cfg.prevLocationsURL = command.Previous
    for _, a := range command.Results {
        fmt.Println(a.Name)
    }
    return nil
}

func commandExpl(cfg *Config, arg ...string) error {
    fmt.Printf("Exploring %s...\n", arg[0])
    area, err := cfg.pokeapiClient.PokemonList(arg[0])
    if err != nil {
       fmt.Println(err)
    }
    fmt.Println("Found Pokemon:")
    for _, encounter := range area.Encounter {
        fmt.Printf(" - %s\n", encounter.Pokemon.Pname)
    }
    return nil
}

type cliCommand struct {
        name        string
        description string
        callback    func(*Config, ...string) error
}

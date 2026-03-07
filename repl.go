package main

import (
    "strings"
    "os"
    "fmt"
    "net/http"
    "io"
    "encoding/json"
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
    }
    return cmds
}

func commandExit() error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandHelp() error {
    list := getCommands()
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println("\n")
    for _, cmd := range list {
        fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    return nil
}

type config struct {
    pokeapiClient    pokeapi.Client
    nextLocationsURL *string
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

func commandMap() error {
    res, err := http.Get("https://pokeapi.co/api/v2/location-area")
    if err != nil {
        return err
    }
    defer res.Body.Close()
    data, err := io.ReadAll(res.Body)
    if err != nil {
        return err
    }
    areas := List{}
    err = json.Unmarshal(data, &areas)
    if err != nil {
        return err
    }
    for _, a := range areas.Results {
        fmt.Println(a.Name)
        config.prevLocationsURL = a.URL
    }
    return nil
}

type cliCommand struct {
        name        string
        description string
        callback    func(*config) error
}

package main

import (
    "strings"
    "os"
    "fmt"
    "github.com/Banarico/pokedex/internal/pokeapi"
    "github.com/Banarico/pokedex/internal/pokecache"
    "math/rand"
    "errors"
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
        "catch": {
            name:        "catch",
            description: "Attempt to catch a given Pokemon",
            callback:    commandCatch,
        },
        "inspect": {
            name:        "inspect",
            description: "View the details of a caught Pokemon",
            callback:    commandInspect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "List all the Pokemon you've caught",
            callback:    commandPkdx,
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
    caughtPkm        map[string]pokeapi.Pokemon
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

func commandCatch(cfg *Config, arg ...string) error {
    fmt.Printf("Throwing a Pokeball at %s...\n", arg[0])
    pkm, err := cfg.pokeapiClient.GetPokemon(arg[0])
    if err != nil {
       return errors.New("That Pokemon doesn't exist!")
    }
    throw := rand.Intn(320)
    if throw >= pkm.BaseExp {
        cfg.caughtPkm[pkm.Name] = pkm
        fmt.Printf("%s was caught!\n", pkm.Name)
    }
    if throw < pkm.BaseExp {
        fmt.Printf("%s escaped!\n", pkm.Name)
    }
    return nil
}

func commandInspect(cfg *Config, arg ...string) error {
    pkm, found := cfg.caughtPkm[arg[0]]
    if found == false {
        return errors.New("You don't have that Pokemon!")
    }
    fmt.Printf("Name: %s\n", pkm.Name)
    fmt.Printf("Height: %d\n", pkm.Height)
    fmt.Printf("Weight: %d\n", pkm.Weight)
    fmt.Println("Stats:")
    for _, s := range pkm.Stats {
        fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
    }
    fmt.Println("Types:")
    for _, t := range pkm.Types {
        fmt.Printf("  - %s\n", t.Type.Name)
    }
    return nil
}

func commandPkdx(cfg *Config, arg ...string) error {
    fmt.Println("Your Pokedex:")
    for _, pkm := range cfg.caughtPkm {
        fmt.Printf(" - %s\n", pkm.Name)
    }
    return nil
}

type cliCommand struct {
        name        string
        description string
        callback    func(*Config, ...string) error
}


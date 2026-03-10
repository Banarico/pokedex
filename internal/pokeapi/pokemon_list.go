package pokeapi

import (
        "encoding/json"
        "io"
        "net/http"
)

type RespAreaPokemon struct {
    Rates []struct {
        Method struct {
            Mname string `json:"name"`
            Murl  string `json:"url"`
        } `json:"encounter_method"`
        Vdetails []struct {
            Rate int `json:"rate"`
            Version struct {
                Vname string `json:"name"`
                Vurl  string `json:"url"`
            } `json:"version"`
        } `json:"version_details"`
    } `json:"encounter_method_rates"`
    Index int `json:"game_index"`
    ID    int `json:"id"`
    Location struct {
        Lname string `json:"name"`
        Lurl  string `json:"url"`
    } `json:"location"`
    Name  string `json:"name"`
    Names []struct {
        Language struct {
            Langname string `json:"name"`
            Langurl  string `json:"url"`
        } `json:"language"`
        Name     string `json:"name"`
    } `json:"names"`
    Encounter []struct {
        Pokemon struct {
            Pname string `json:"name"`
            Purl  string `json:"url"`
        } `json:"pokemon"`
        Pdetails []struct {
            Edetails []struct {
                Chance     int `json:"chance"`
                Conditions []struct {
                }`json:"condition_values"`
                Mxlevel    int `json:"max_level"`
                Emethod    struct {
                    Ename string `json:"name"`
                    Eurl  string `json:"url"`
                } `json:"method"`
                Minlevel   int `json:"min_level"`
            } `json:"encounter_details"`
            Mxchance int `json:"max_chance"`
            Pversion struct {
                Gname string `json:"name"`
                Gurl  string `json:"url"`
            } `json:"version"`
        } `json:"version_details"`
    } `json:"pokemon_encounters"`
}

func (c *Client) PokemonList(arg string) (RespAreaPokemon, error) {
    url := "https://pokeapi.co/api/v2/location-area/" + arg
    if cached, ok := c.cache.Get(url); ok {
        pokemonResp := RespAreaPokemon{}
        err := json.Unmarshal(cached, &pokemonResp)
        if err != nil {
            return RespAreaPokemon{}, err
        }
        return pokemonResp, nil
    }
    req, err := http.Get(url)
    if err != nil {
        return RespAreaPokemon{}, err
    }
    data, err := io.ReadAll(req.Body)
    if err != nil {
        return RespAreaPokemon{}, err
    }
    pokemonResp := RespAreaPokemon{}
    err = json.Unmarshal(data, &pokemonResp)
    if err != nil {
        return RespAreaPokemon{}, err
    }
    c.cache.Add(url, data)
    return pokemonResp, nil
}

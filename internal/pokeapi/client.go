package pokeapi

import (
    "time"
    "net/http"
    "github.com/Banarico/pokedex/internal/pokecache"
)

type Client struct {
    cache      pokecache.Cache
    httpClient http.Client
}

func NewClient(timeout time.Duration, cache pokecache.Cache) Client {
    return Client{
        cache: cache,
        httpClient: http.Client{
            Timeout: timeout,
        },
    }
}

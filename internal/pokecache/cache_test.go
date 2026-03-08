package pokecache

import (
    "testing"
    "fmt"
    "time"
)

func TestAddGet(t *testing.T) {
    const interval = 5 * time.Second
    cases := []struct {
        key string
        val []byte
    }{
        {
            key: "https://pokeapi.co/api/v2/location-area",
            val: []byte("first_page_data"),
        },
        {
            key: "https://pokeapi.co/api/v2/location-area?offset=20&limit=20",
            val: []byte("second_page_data"),
        },
    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
            cache := NewCache(interval)
            
            // 1. Add the "URL" and "JSON data" to the cache
            cache.Add(c.key, c.val)
            
            // 2. Try to retrieve it
            val, ok := cache.Get(c.key)
            if !ok {
                t.Errorf("expected to find key %s", c.key)
                return
            }
            
            // 3. Check if the data matches
            if string(val) != string(c.val) {
                t.Errorf("expected to find value %s, got %s", string(c.val), string(val))
                return
            }
        })
    }
}

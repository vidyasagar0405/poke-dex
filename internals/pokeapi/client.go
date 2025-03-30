package pokeapi

import (
	"net/http"
	"time"

	"github.com/vidyasagar0405/pokedexcli/internals/pokecache"
)

type Client struct {
	httpClient http.Client
    Cache pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
        Cache: *pokecache.NewCache(5 * time.Second),
	}
}


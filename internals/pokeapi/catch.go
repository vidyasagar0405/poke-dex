package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemon string) (Pokemon, error) {
    url := BaseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		pokeInfoResp := Pokemon{}
		err := json.Unmarshal(val, &pokeInfoResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokeInfoResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad StatusCode: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokeInfoResp := Pokemon{}
	err = json.Unmarshal(data, &pokeInfoResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokeInfoResp, nil
}

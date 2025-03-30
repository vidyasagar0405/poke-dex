package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocationArea(locArea string) (locationExploreResp, error) {
    url := BaseURL + "/location-area/" + locArea

	if val, ok := c.cache.Get(url); ok {
		exploreLocationResp := locationExploreResp{}
		err := json.Unmarshal(val, &exploreLocationResp)
		if err != nil {
			return locationExploreResp{}, err
		}
		return exploreLocationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationExploreResp{}, err
	}

    fmt.Println("Exploring pastoria-city-area...")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationExploreResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationExploreResp{}, fmt.Errorf("bad StatusCode: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationExploreResp{}, err
	}

	exploreLocationResp := locationExploreResp{}
	err = json.Unmarshal(data, &exploreLocationResp)
	if err != nil {
		return locationExploreResp{}, err
	}

	c.cache.Add(url, data)

	return exploreLocationResp, nil
}

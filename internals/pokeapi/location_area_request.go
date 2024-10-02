package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas() (locationAreasResp, error) {
	endpoint := "/location-area"
	fullUrl := baseURL + endpoint

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return locationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return locationAreasResp{}, fmt.Errorf("bad StatusCode: %v", resp.StatusCode)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreasResp{}, err
	}

	location_areas_resp := locationAreasResp{}
	err = json.Unmarshal(data, &location_areas_resp)
	if err != nil {
		return locationAreasResp{}, err
	}

	return location_areas_resp, nil
}

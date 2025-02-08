package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

func (c *Client) ListLocationAreas(url string) (locationAreasResp, error) {

    req, err := http.NewRequest("GET", url, nil)
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

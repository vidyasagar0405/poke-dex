package pokeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	timeout := 10 * time.Second
	client := NewClient(timeout)

	if client.httpClient.Timeout != timeout {
		t.Errorf("Expected timeout to be %v, got %v", timeout, client.httpClient.Timeout)
	}

	// Since Cache is a value type, we can't check if it's nil
	// Instead, we can check if it's initialized by checking internal state
	if reflect.ValueOf(client.Cache).IsZero() {
		t.Error("Expected cache to be initialized")
	}
}

func TestListLocationAreas_Success(t *testing.T) {
	// Create a mock server
	mockResp := locationAreasResp{
		Count: 2,
		Next:  stringPtr("https://pokeapi.co/api/v2/location-area?offset=20&limit=20"),
		Prev:  nil,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{
				Name: "area1",
				URL:  "https://pokeapi.co/api/v2/location-area/1/",
			},
			{
				Name: "area2",
				URL:  "https://pokeapi.co/api/v2/location-area/2/",
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer mockServer.Close()

	// Create a client
	client := NewClient(10 * time.Second)

	// Call the method
	resp, err := client.ListLocationAreas(mockServer.URL)

	// Check the result
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.Count != mockResp.Count {
		t.Errorf("Expected Count to be %d, got %d", mockResp.Count, resp.Count)
	}

	if *resp.Next != *mockResp.Next {
		t.Errorf("Expected Next to be %s, got %s", *mockResp.Next, *resp.Next)
	}

	if resp.Prev != mockResp.Prev {
		t.Errorf("Expected Prev to be %v, got %v", mockResp.Prev, resp.Prev)
	}

	if !reflect.DeepEqual(resp.Results, mockResp.Results) {
		t.Errorf("Expected Results to be %v, got %v", mockResp.Results, resp.Results)
	}
}

func TestListLocationAreas_BadStatusCode(t *testing.T) {
	// Create a mock server that returns a 404
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer mockServer.Close()

	// Create a client
	client := NewClient(10 * time.Second)

	// Call the method
	_, err := client.ListLocationAreas(mockServer.URL)

	// Check the error
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}

func TestListLocationAreas_InvalidJSON(t *testing.T) {
	// Create a mock server that returns invalid JSON
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("invalid json"))
	}))
	defer mockServer.Close()

	// Create a client
	client := NewClient(10 * time.Second)

	// Call the method
	_, err := client.ListLocationAreas(mockServer.URL)

	// Check the error
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}

func TestListLocationAreas_ServerError(t *testing.T) {
	// Create a mock server that returns a 500
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	// Create a client
	client := NewClient(10 * time.Second)

	// Call the method
	_, err := client.ListLocationAreas(mockServer.URL)

	// Check the error
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}

func TestListLocationAreas_CacheStorage(t *testing.T) {
	// Create a mock server
	mockResp := locationAreasResp{
		Count: 1,
		Results: []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		}{
			{
				Name: "test-area",
				URL:  "https://pokeapi.co/api/v2/location-area/1/",
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer mockServer.Close()

	// Create a client with a longer cache duration for testing
	client := NewClient(10 * time.Second)
	
	// Make a request
	_, err := client.ListLocationAreas(mockServer.URL)
	if err != nil {
		t.Fatalf("Expected no error on request, got %v", err)
	}
	
	// Check that the cache contains the data
	cachedData, exists := client.Cache.Get(mockServer.URL)
	if !exists {
		t.Error("Expected data to be in the cache, but it wasn't found")
	}
	if cachedData == nil {
		t.Error("Expected cached data to not be nil")
	}
	
	// Verify we can decode the cached data
	var result locationAreasResp
	err = json.Unmarshal(cachedData, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal cached data: %v", err)
	}
	
	// Verify the cached data matches our expected response
	if result.Count != mockResp.Count {
		t.Errorf("Expected cached Count to be %d, got %d", mockResp.Count, result.Count)
	}
	
	if !reflect.DeepEqual(result.Results, mockResp.Results) {
		t.Errorf("Expected cached Results to match original response")
	}
}

// Helper function to create string pointers
func stringPtr(s string) *string {
	return &s
}

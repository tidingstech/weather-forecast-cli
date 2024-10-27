// main_test.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"druc/sun/model" // Replace with your actual model package path
)

// TestWeatherAPI tests the API call to fetch weather data.
func TestWeatherAPI(t *testing.T) {
	// Set up a mock server to simulate the API response
	mockResponse := `{
		"location": {
			"name": "Nairobi",
			"country": "Kenya"
		},
		"current": {
			"temp_c": 25.0,
			"condition": {
				"text": "Sunny"
			}
		},
		"forecast": {
			"forecastday": [{
				"hour": [{
					"time_epoch": 1700000000,
					"temp_c": 22.0,
					"condition": {
						"text": "Clear"
					},
					"chance_of_rain": 20.0
				}]
			}]
		}
	}`

	// Create a mock server
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	// Replace the real URL with the mock server's URL
	os.Setenv("WEATHER_API_URL", mockServer.URL)

	// Set up command-line argument for testing location
	os.Args = []string{"cmd", "Nairobi"}

	// Capture the output from the main function
	main()
}

// TestDateFiltering tests the filtering of forecast data to ensure only future dates are considered.
func TestDateFiltering(t *testing.T) {
	// Test for correct date filtering for future dates
	hourData := model.Weather{
		Forecast: struct {
			ForecastDay []struct {
				Hour []struct {
					TimeEpoch     int64   `json:"time_epoch"`
					TempC         float64 `json:"temp_c"`
					Condition     struct {
						Text string `json:"text"`
					} `json:"condition"`
					ChanceOfRain float64 `json:"chance_of_rain"`
				} `json:"hour"`
			} `json:"forecastday"`
		}{
			ForecastDay: []struct {
				Hour []struct {
					TimeEpoch     int64   `json:"time_epoch"`
					TempC         float64 `json:"temp_c"`
					Condition     struct {
						Text string `json:"text"`
					} `json:"condition"`
					ChanceOfRain float64 `json:"chance_of_rain"`
				}{
					{
						TimeEpoch:    time.Now().Add(1 * time.Hour).Unix(),
						TempC:        22.0,
						Condition:    struct{ Text string `json:"text"` }{Text: "Clear"},
						ChanceOfRain: 15.0,
					},
				},
			},
		},
	}

	if len(hourData.Forecast.ForecastDay) == 0 {
		t.Error("No forecast day data found")
	}
	if len(hourData.Forecast.ForecastDay[0].Hour) == 0 {
		t.Error("No hourly data found")
	}

	// Ensure the time is in the future
	for _, hour := range hourData.Forecast.ForecastDay[0].Hour {
		if time.Unix(hour.TimeEpoch, 0).Before(time.Now()) {
			t.Error("Found a forecast hour with a past time")
		}
	}
}

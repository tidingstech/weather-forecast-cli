package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"

    "github.com/fatih/color"
    "druc/sun/model" // Replace with the actual path to your model package
)

func main() {
    // Set default location and allow overriding via command-line arguments
    q := "Kenya"
    if len(os.Args) >= 2 {
        q = os.Args[1]
    }

    // Fetch weather data from the API
    res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=3caac56eb5e74ef7b82145754242710&q=" + q + "&days=1&aqi=no&alerts=no")
    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    // Check if the response status is 200 (OK)
    if res.StatusCode != http.StatusOK {
        panic("Weather API not available")
    }

    // Read the response body
    body, err := io.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }

    // Unmarshal JSON data into the Weather struct from the model package
    var weather model.Weather
    err = json.Unmarshal(body, &weather)
    if err != nil {
        panic(err)
    }

    // Print the current weather information
    location, current := weather.Location, weather.Current
    fmt.Printf(
        "%s, %s: %.0f°C, %s\n",
        location.Name,
        location.Country,
        current.TempC,
        current.Condition.Text,
    )

    // Ensure forecast data is available
    if len(weather.Forecast.ForecastDay) > 0 && len(weather.Forecast.ForecastDay[0].Hour) > 0 {
        // Print the forecast details for each hour
        for _, hour := range weather.Forecast.ForecastDay[0].Hour {
            date := time.Unix(hour.TimeEpoch, 0)
            if date.Before(time.Now()) {
                continue
            }

            // Format the message
            message := fmt.Sprintf(
                "%s - %.0f°C, %.0f%% chance of rain, %s\n",
                date.Format("15:04"),
                hour.TempC,
                hour.ChanceOfRain,
                hour.Condition.Text,
            )

            // Print message based on rain chance
            if hour.ChanceOfRain < 40 {
                color.Green(message) // Green indicates favorable conditions
            } else {
                color.Red(message) // Red indicates caution or higher rain chances
            }
        }
    } else {
        fmt.Println("No forecast data available.")
    }
}

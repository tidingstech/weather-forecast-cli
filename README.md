# Weather Forecast CLI Application

This is a simple Go-based CLI application that fetches weather data from an external API and displays current weather conditions along with forecast information. The project uses Go's built-in packages and imports a custom model to parse JSON responses from the WeatherAPI.

## Features
- Fetch current weather information based on the provided location.
- Display temperature, weather conditions, and chance of rain for the specified location.
- Highlights rain chances of more than 40% using colored output.
- Uses a command-line argument to specify a location (defaults to "Kenya" if not provided).

## Prerequisites

- Go 1.16 or later
- [WeatherAPI](https://www.weatherapi.com/) key (for making API calls)
- Internet connection

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/Oragwel/weather-forecast-cli.git
    cd weather-forecast-cli
    ```

2. **Setup the project**:
    Ensure that you have a valid API key from [WeatherAPI](https://www.weatherapi.com/). You can get a free API key by signing up.

3. **Set up the `GOPATH`**:
    Make sure the `GOPATH` is properly set on your system.

## Project Structure

```
weather-forecast-cli/
│
├── main.go                  # Main application code
├── model/
│   └── model.go             # Contains the Weather struct to parse the JSON response
├── main_test.go             # Test cases for the application
├── go.mod                   # Go module file
└── README.md                # Project documentation
```

## Usage

1. **Run the application**:

    ```bash
    go run main.go [location]
    ```

    For example:

    ```bash
    go run main.go Nairobi
    ```

    If no location is provided, it defaults to "Kenya".

2. **Build the application**:

    ```bash
    go build -o weather-cli main.go
    ```

3. **Execute the built application**:

    ```bash
    ./weather-cli Nairobi
    ```

4. **Run the tests**:

    ```bash
    go test -v
    ```

## Dependencies

- [fatih/color](https://github.com/fatih/color) - For colored output in the terminal.
- Standard Go libraries such as `net/http`, `encoding/json`, `os`, and `time`.


## Environment Variables

Set the following environment variables:

- `WEATHER_API_URL` (optional): If you are using a different API endpoint for testing or production.
- `WEATHER_API_KEY`: Your WeatherAPI key.

## Example Output

```text
Nairobi, Kenya: 25C, Sunny
15:00 - 22°C, 20% chance of rain, Clear
16:00 - 21°C, 10% chance of rain, Clear
```

## Contributing

If you find a bug or have a suggestion for improvements, please feel free to open an issue or submit a pull request. Contributions are always welcome!

## Contributor

- **Otieno Ragwel Rogers**

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
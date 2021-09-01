package cmd

import "os"

type Config struct {
	OpenWeather  string
	CustomSearch string
	EngineID     string
}

func LoadConfig() Config {
	return Config{
		OpenWeather:  os.Getenv("OPEN_WEATHER_API_KEY"),
		CustomSearch: os.Getenv("GOOGLE_SEARCH_API_KEY"),
		EngineID:     os.Getenv("GOOGLE_SEARCH_ENGINE_ID"),
	}
}

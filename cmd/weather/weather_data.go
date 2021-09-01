package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ocakhasan/ocak/cmd"
)

type weatherInfo struct {
	Main        string
	Description string
}

type main struct {
	Temp    float32 `json:"temp"`
	Feels   float32 `json:"feels_like"`
	TempMin float32 `json:"temp_min"`
	TempMax float32 `json:"temp_max"`
}

type WeatherData struct {
	Weather []*weatherInfo `json:"weather"`
	Main    main           `json:"main"`
	Name    string         `json:"name"`
}

func GetWeatherData(city string) error {
	config := cmd.LoadConfig()
	if config.OpenWeather == "" {
		return fmt.Errorf("API_KEY not found. You need to specify API_KEY as your environment variable")
	}
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, config.OpenWeather)
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("there is an error fetching openweathermap %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status Code is %d for url %s", resp.StatusCode, url)
	}
	var data WeatherData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return fmt.Errorf("there is error in decoding data: %v", err)
	}

	if err := Report.Execute(os.Stdout, data); err != nil {
		return fmt.Errorf("template error: %v", err)
	}
	return nil
}

package forecast

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Forecast struct {
	City string `json:"city_name"`
	Data []Day  `json:"data"`
}

type Day struct {
	Date       string  `json:"valid_date"`
	MinTemp    float64 `json:"min_temp"`
	MaxTemp    float64 `json:"max_temp"`
	PrecipProb int     `json:"pop"`
	Precip     float64 `json:"precip"`
}

var client *http.Client

func GetForecast(zip string) (Forecast, error) {
	var forecast Forecast
	client = &http.Client{Timeout: 10 * time.Second}

	envErr := godotenv.Load(".env")
	if envErr != nil {
		return forecast, envErr
	}

	apiKey := os.Getenv("API_KEY")

	resp, apiErr := client.Get("https://api.weatherbit.io/v2.0/forecast/daily?units=I&postal_code=" + zip + "&key=" + apiKey)
	if apiErr != nil {
		return forecast, apiErr
	}
	if resp.StatusCode != 200 {
		return forecast, errors.New("API Error")
	}

	defer resp.Body.Close()

	decodeErr := json.NewDecoder(resp.Body).Decode(&forecast)
	if decodeErr != nil {
		return forecast, decodeErr
	}

	return forecast, nil
}

func FormatForecast(f Forecast) string {
	result := fmt.Sprintf("10-Day Forecast - %s", f.City)

	for _, day := range f.Data[0:10] {
		dateSlice := strings.Split(day.Date, "-")
		year := dateSlice[0]
		month := dateSlice[1]
		dayInt := dateSlice[2]
		formattedDate := fmt.Sprintf("%v-%v-%v", month, dayInt, year)

		result += fmt.Sprintf(`
  %s
    Min Temp: %.1f°F
    Max Temp: %.1f°F
    Probability of Precipitation: %d%%
    Precipitation: %.1f inches`,
			formattedDate,
			day.MinTemp,
			day.MaxTemp,
			day.PrecipProb,
			day.Precip,
		)
	}

	return result
}

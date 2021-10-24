package gateway

import (
	"errors"
	"net/http"

	"github.com/imroc/req"
)

var (
	ErrOpenWeatherMapWrongStatusCode = errors.New("gateway_openweathermap: wrong status code")
	ErrOpenWeatherMapCanNotParseBody = errors.New("gateway_openweathermap: can not parse body")
)

type OpenWeatherMapResponse struct {
	Main OpenWeatherMapMain `json:"main"`
	Wind OpenWeatherMapWind `json:"wind"`
}

type OpenWeatherMapMain struct {
	Temp float32 `json:"temp"`
}
type OpenWeatherMapWind struct {
	Speed float32 `json:"speed"`
}

func (g *gateway) GetInfoFromOpenWeatherMap(city string) (WeatherDTO, error) {

	var weatherDTO WeatherDTO
	var response OpenWeatherMapResponse

	params := req.Param{
		"q":     city,
		"appid": g.cfg.GetString("OPENWEATHERMAP_TOKEN"),
	}

	resp, err := req.Get(g.cfg.GetString("OPENWEATHERMAP_ADDRESS"), params)

	if err != nil {
		return weatherDTO, err
	}

	if resp.Response().StatusCode != http.StatusOK {
		return weatherDTO, ErrOpenWeatherMapWrongStatusCode
	}

	err = resp.ToJSON(&response)

	if err != nil {
		return weatherDTO, err
	}

	weatherDTO.Temperature = response.Main.Temp
	weatherDTO.WindSpeed = response.Wind.Speed
	return weatherDTO, nil
}

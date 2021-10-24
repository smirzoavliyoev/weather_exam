package gateway

import (
	"errors"

	"github.com/imroc/req"
)

var (
	ErrWeatherStackNotSuccessError = errors.New("gateway_weatherstack: response is not success")
	ErrWrongReponseFormat          = errors.New("gateway_weatherstack: wrong response format")
)

type WeatherstackFaultResponse struct {
	Success string                    `json:"success"`
	Error   WeatherstackErrorResponse `json:"error"`
}

type WeatherstackErrorResponse struct {
	Code int `json:"code"`
	//todo use enum and validation
	Type string `json:"type"`
	Info string `json:"info"`
}

type WeatherstackSuccessResponse struct {
	Current WeatherstackCurrentState `json:"current"`
}

type WeatherstackCurrentState struct {
	Temperature float32 `json:"temperature"`
	WindSpeed   float32 `json:"wind_speed"`
}

func (g *gateway) GetInfoFromWeatherstack(city string) (WeatherDTO, error) {
	var weatherstackDTO WeatherDTO

	param := req.Param{
		"access_key": g.cfg.GetString("WEATHERSTACK_TOKEN"),
		"query":      city,
	}

	resp, err := req.Get(g.cfg.GetString("WEATHERSTACK_ADDRESS"), param)
	if err != nil {
		return weatherstackDTO, err
	}

	var response interface{}

	err = resp.ToJSON(&response)

	if err != nil {
		return weatherstackDTO, err
	}

	res, ok := response.(WeatherstackSuccessResponse)
	if !ok {
		return weatherstackDTO, ErrWeatherStackNotSuccessError
	}

	weatherstackDTO.Temperature = res.Current.Temperature
	weatherstackDTO.WindSpeed = res.Current.WindSpeed

	return weatherstackDTO, nil
}

package gateway

import "weather/pkg/config"

type WeatherDTO struct {
	WindSpeed   float32
	Temperature float32
}
type Gateway interface {
	GetInfoFromWeatherstack(city string) (WeatherDTO, error)
	GetInfoFromOpenWeatherMap(city string) (WeatherDTO, error)
}

type gateway struct {
	cfg config.Configer
}

func NewGateway(cfg config.Configer) Gateway {
	return &gateway{
		cfg: cfg,
	}
}

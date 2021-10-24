package internal

import (
	"context"
	"strings"
	"time"
	"weather/pkg/cache"
	"weather/pkg/gateway"
	"weather/pkg/logger"
	"weather/pkg/structs"

	"go.uber.org/zap"
)

var City string = "Sidney"
var city string = "sidney"

func CacheWeatherResults(gateway gateway.Gateway, cache cache.Cacher, logger logger.Logger) error {

	//documentation shows that there can be uppercase or lowecase letters
	weatherDto, err := gateway.GetInfoFromWeatherstack(City)

	if err == nil {

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		errCacheSave := cache.Set(ctx, "weatherInfo."+strings.ToLower(City),
			structs.WeatherInfo{
				Temperature: weatherDto.Temperature,
				WindSpeedd:  weatherDto.WindSpeed,
			},
			5*time.Second)
		if errCacheSave != nil {
			logger.Error("cache: can not save", zap.Error(errCacheSave))
			return errCacheSave
		}

		logger.Info("data cached", zap.Any("some1", weatherDto))
		return nil
	} else {
		logger.Error("error from weatherstack service", zap.Error(err))
	}

	weatherDto, err = gateway.GetInfoFromOpenWeatherMap(city)

	if err != nil {
		logger.Error("can not get data from openweathermap", zap.Error(err))
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	err = cache.Set(ctx, "weatherInfo."+strings.ToLower(city), structs.WeatherInfo{
		Temperature: weatherDto.Temperature - 273.15,
		WindSpeedd:  weatherDto.WindSpeed,
	},
		5*time.Second)

	if err != nil {
		logger.Error("can not save to cache", zap.Error(err))
		return err
	}

	logger.Info("data cached", zap.Any("some", weatherDto))
	return nil
}

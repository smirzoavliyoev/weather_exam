package weatherinfo

import (
	"context"
	"strings"
	"time"
	"weather/pkg/cache"
	"weather/pkg/logger"
	"weather/pkg/structs"

	"go.uber.org/zap"
)

type WeatherInfoProvider interface {
	GetWeatrherInfo(string) (structs.WeatherInfo, error)
}

type weatherinfoprovider struct {
	cache  cache.Cacher
	logger logger.Logger
}

func NewWeatherInfoProvider(cache cache.Cacher, logger logger.Logger) WeatherInfoProvider {
	return &weatherinfoprovider{
		cache:  cache,
		logger: logger,
	}
}

func (w *weatherinfoprovider) GetWeatrherInfo(city string) (structs.WeatherInfo, error) {
	var weatherInfo structs.WeatherInfo

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// todo: performance review - useless marshaling and unmarshaling that can influence on performance
	// in this case when nothing should do with data
	// in cases when have to opperate with them its ok
	err := w.cache.Get(ctx, "weatherInfo."+strings.ToLower(city), &weatherInfo)
	if err != nil {
		w.logger.Error("can not get weather info", zap.Error(err))
		return weatherInfo, err
	}

	return weatherInfo, nil
}

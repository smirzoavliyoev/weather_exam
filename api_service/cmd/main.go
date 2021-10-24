package main

import (
	"weather/api_service/cmd/handler"
	"weather/api_service/cmd/router"
	"weather/api_service/internal/weatherinfo"
	"weather/pkg/cache"
	"weather/pkg/config"
	"weather/pkg/logger"
	"weather/pkg/tracer"
)

func main() {

	configer := config.NewConfig()
	logger := logger.NewLogger()
	cacher := cache.GetCacher(configer)
	weatherInfoProvider := weatherinfo.NewWeatherInfoProvider(cacher, logger)
	tracer := tracer.NewTracer(configer, logger)
	handler := handler.GetHandler(configer, cacher, weatherInfoProvider, logger, tracer)

	router.Handle(configer, handler, logger)
}

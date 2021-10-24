package main

import (
	"fmt"
	"weather/pkg/cache"
	"weather/pkg/config"
	"weather/pkg/gateway"
	log "weather/pkg/logger"
	"weather/worker_service/internal"

	"github.com/jasonlvhit/gocron"
)

func main() {
	cfg := config.NewConfig()
	logger := log.NewLogger()
	cache := cache.GetCacher(cfg)
	gateway := gateway.NewGateway(cfg)

	gocron.Every(3).Second().Do(internal.CacheWeatherResults, gateway, cache, logger)
	<-gocron.Start()

	err := log.Sync()
	if err != nil {
		fmt.Println(err)
	}
}

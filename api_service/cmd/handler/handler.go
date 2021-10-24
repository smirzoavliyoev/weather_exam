package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"weather/api_service/internal/weatherinfo"
	"weather/pkg/cache"
	"weather/pkg/config"
	"weather/pkg/logger"
	"weather/pkg/responser"
	"weather/pkg/tracer"
)

type Handler interface {
	GetWeaherStatus(w http.ResponseWriter, r *http.Request)
}

func GetHandler(cfg config.Configer,
	cache cache.Cacher,
	WeatherInfoProvider weatherinfo.WeatherInfoProvider,
	logger logger.Logger,
	tracer tracer.Tracer) Handler {
	return &handler{
		cfg:                 cfg,
		cache:               cache,
		weatherinfoProvider: WeatherInfoProvider,
		logger:              logger,
		tracer:              tracer,
	}
}

type handler struct {
	cfg                 config.Configer
	cache               cache.Cacher
	weatherinfoProvider weatherinfo.WeatherInfoProvider
	logger              logger.Logger
	tracer              tracer.Tracer
}

func (h *handler) GetWeaherStatus(w http.ResponseWriter, r *http.Request) {
	var (
		params = r.URL.Query()
	)
	span, _ := h.tracer.StartSpanFromContext(context.Background(), "GetWeatherStatus")
	defer span.Finish()

	city, ok := params["city"]
	if !ok {
		responser.Response(responser.BadRequest, w)
		return
	}

	if len(city) != 1 {
		responser.Response(responser.BadRequest, w)
		return
	}

	weatherInfo, err := h.weatherinfoProvider.GetWeatrherInfo(city[0])

	if err != nil {
		responser.Response(responser.InternalError, w)
		return
	}

	body, err := json.Marshal(weatherInfo)
	if err != nil {
		responser.Response(responser.InternalError, w)
		return
	}

	io.WriteString(w, string(body))
}

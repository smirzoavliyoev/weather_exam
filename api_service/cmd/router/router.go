package router

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"weather/api_service/cmd/handler"
	"weather/pkg/config"
	log "weather/pkg/logger"

	"go.uber.org/zap"
)

func Handle(cfg config.Configer, handler handler.Handler, logger log.Logger) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			w.Write([]byte("method not allowed"))
		}

		io.WriteString(w, "ok")
	})

	http.HandleFunc("/"+cfg.GetString("APP_VERSION")+"/weather", handler.GetWeaherStatus)

	var server *http.Server = &http.Server{
		Addr:        cfg.GetString("SERVER_PORT"),
		IdleTimeout: 1 * time.Second,
	}

	var appStart = func(cfg config.Configer, server *http.Server) error {
		err := server.ListenAndServe()
		if err != nil {
			logger.Error("can not start server", zap.Error(err))
			return err
		}
		return nil
	}

	var appStop = func(ctx context.Context, server *http.Server) error {
		fmt.Println("Gracefully shutdown")
		err := server.Shutdown(ctx)
		if err != nil {
			logger.Error("can not shutdown server", zap.Error(err))
			return err
		}
		err = log.Sync()
		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}

	term := make(chan os.Signal, 1)

	signal.Notify(term, syscall.SIGTERM, syscall.SIGINT)

	go appStart(cfg, server)

	<-term
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	appStop(ctx, server)

}

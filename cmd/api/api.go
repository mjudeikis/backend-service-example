package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

	"github.com/mjudeikis/backend-service/pkg/service"
	"github.com/mjudeikis/backend-service/pkg/utils/logger"
	"go.uber.org/zap"
)

var (
	logLevel = flag.String("log-level", "info", `LogLevel`)
)

func main() {
	flag.Parse()
	ctx := context.Background()

	log := logger.GetLoggerInstance("", zap.DebugLevel)

	log.Info("starting api")

	if err := run(ctx, log); err != nil {
		log.Error("error starting controller", zap.Error(err))
	}
}

func run(ctx context.Context, log *zap.Logger) error {

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	stop := make(chan struct{})
	done := make(chan struct{})

	service, err := service.New(log)
	if err != nil {
		return err
	}

	go service.Run(ctx, stop, done)
	<-signals
	// we catch both sigterm (used by systemd) and Interupt (ctr+c) for development. Later is not really needed
	log.Info("received Sigterm/Int")
	close(stop)

	shutdownCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()
	select {
	case <-shutdownCtx.Done():
		log.Warn("service didn't shutdown in time, force exit")
	case <-done:
		// OK
	}

	return nil

}

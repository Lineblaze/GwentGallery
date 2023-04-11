package main

import (
	"context"
	"github.com/Lineblaze/GwentGallery/app/internal/app"
	"github.com/Lineblaze/GwentGallery/app/internal/infrastructure/config"
	"github.com/Lineblaze/GwentGallery/app/pkg/logging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logging.Info(ctx, "config initializing")
	cfg := config.GetConfig()

	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		logging.Fatal(ctx, err)
	}

	logging.Info(ctx, "Running Application")
	a.Run(ctx)
}

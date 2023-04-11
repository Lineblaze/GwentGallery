package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/Lineblaze/GwentGallery/app/internal/infrastructure/config"
	"github.com/Lineblaze/GwentGallery/app/pkg/client/postgresql"
	"github.com/Lineblaze/GwentGallery/app/pkg/logging"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net"
	"net/http"
	"time"
)

type App struct {
	cfg        *config.Config
	router     *httprouter.Router
	httpServer *http.Server
	pgClient   *pgxpool.Pool
}

func NewApp(ctx context.Context, config *config.Config) (App, error) {
	logging.Info(ctx, "router initializing")
	router := httprouter.New()

	logging.Info(ctx, "swagger docs initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)
	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}

	return App{
		cfg:      config,
		router:   router,
		pgClient: pgClient,
	}, nil
}

func (a *App) Run(ctx context.Context) {
	a.startHTTP(ctx)
}

func (a *App) startHTTP(ctx context.Context) error {
	logger := logging.WithFields(ctx, map[string]interface{}{
		"IP":   a.cfg.HTTP.IP,
		"Port": a.cfg.HTTP.Port,
	})
	logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		logger.WithError(err).Fatal("failed to create listener")
	}

	logger.WithFields(map[string]interface{}{
		"AllowedMethods":     a.cfg.HTTP.CORS.AllowedMethods,
		"AllowedOrigins":     a.cfg.HTTP.CORS.AllowedOrigins,
		"AllowCredentials":   a.cfg.HTTP.CORS.AllowCredentials,
		"AllowedHeaders":     a.cfg.HTTP.CORS.AllowedHeaders,
		"OptionsPassthrough": a.cfg.HTTP.CORS.OptionsPassthrough,
		"ExposedHeaders":     a.cfg.HTTP.CORS.ExposedHeaders,
		"Debug":              a.cfg.HTTP.CORS.Debug,
	})
	c := cors.New(cors.Options{
		AllowedMethods:     a.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins:     a.cfg.HTTP.CORS.AllowedOrigins,
		AllowCredentials:   a.cfg.HTTP.CORS.AllowCredentials,
		AllowedHeaders:     a.cfg.HTTP.CORS.AllowedHeaders,
		OptionsPassthrough: a.cfg.HTTP.CORS.OptionsPassthrough,
		ExposedHeaders:     a.cfg.HTTP.CORS.ExposedHeaders,
		Debug:              a.cfg.HTTP.CORS.Debug,
	})

	handler := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warning("server shutdown")
		default:
			logger.Fatal(err)
		}
	}
	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		logger.Fatal(err)
	}
	return err
}

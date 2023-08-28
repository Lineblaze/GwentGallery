package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/Lineblaze/GwentGallery/app/internal/core"
	"github.com/Lineblaze/GwentGallery/app/internal/infrastructure/config"
	"github.com/Lineblaze/GwentGallery/app/internal/infrastructure/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Server interface {
	Run() error
	Shutdown(ctx context.Context) error
}

type server struct {
	core       core.Core
	router     chi.Router
	httpServer *http.Server
}

func NewHttpServer(core core.Core) (Server, error) {
	router := chi.NewRouter()
	httpPort := config.Config.Interfaces.Http.Port
	httpHost := config.Config.Interfaces.Http.Host

	s := &server{
		core:   core,
		router: router,
		httpServer: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", httpHost, httpPort),
			Handler: router,
		},
	}

	s.initRouter()

	return s, nil
}

func (s *server) Run() error {
	if s.httpServer != nil {
		go func() {
			logger.Logger.Info().Msgf("Starting HTTP server at %s", s.httpServer.Addr)
			err := s.httpServer.ListenAndServe()
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				logger.Logger.Err(err)
			}
		}()
	}

	logger.Logger.Info().Msgf("Starting HTTP server at %s", s.httpServer.Addr)
	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

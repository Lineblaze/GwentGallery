package http

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"os"
)

func (s *server) initRouter() http.Handler {
	r := s.router

	if os.Getenv("ENV") == "dev" {
		r.Use(middleware.Logger)
	}
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	r.Post("/api/v1/auth/sign-in", s.SignIn)
	r.Post(`/api/v1/auth/sign-up`, s.SignUp)

	r.Get(`/api/v1/users/{user_id}`, s.GetUser)

	r.Get(`/api/v1/cards/{card_id}`, s.GetCard)
	r.Post(`/api/v1/cards`, s.CreateCard)
	r.Put(`/api/v1/cards/{card_id}`, s.UpdateCard)
	r.Delete(`/api/v1/cards/{card_id}`, s.DeleteCard)

	return r
}

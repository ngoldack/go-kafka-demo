package http

import (
	"go-kafka-demo/services/gateway/internal/adapter/config"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	r *chi.Mux
}

func NewRouter(cfg config.HttpConfig) (*Router, error) {
	r := &Router{
		r: chi.NewRouter(),
	}

	return r, nil
}

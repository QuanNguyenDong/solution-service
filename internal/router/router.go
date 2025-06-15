package router

import (
	"github.com/QuanNguyenDong/solution-service/internal/handler"
	"github.com/go-chi/chi/v5"
)

func New() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/healthz", handler.Healthz)
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/", handler.Home)
	})
	return router
}

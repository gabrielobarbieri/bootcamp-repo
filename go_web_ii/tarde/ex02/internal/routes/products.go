package routes

import (
	"github.com/gabrielobarbieri/api-meli/internal/controller"
	"github.com/gabrielobarbieri/api-meli/internal/middleware"

	"github.com/go-chi/chi"
)

func RegisterProductRoutes(r chi.Router, productCtrl *controller.HandlerProducts) {
	r.Route("/products", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Get("/ping", productCtrl.Ping)
		r.Post("/", productCtrl.Create)
		r.Get("/", productCtrl.Get)
		r.Get("/{id}", productCtrl.GetById)
		r.Get("/search", productCtrl.Search)
		r.Put("/{id}", productCtrl.UpdateOrCreate)
		r.Patch("/{id}", productCtrl.Update)
		r.Delete("/{id}", productCtrl.Delete)
	})
}

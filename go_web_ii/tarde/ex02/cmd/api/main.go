package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielobarbieri/api-meli/cmd/handlers"
	"github.com/gabrielobarbieri/api-meli/internal/storage"
	"github.com/gabrielobarbieri/api-meli/internal/storage/loader"

	"github.com/go-chi/chi/v5"
)

func main() {
	ld := loader.NewLoaderJSON("../../docs/db/json/products.json")
	DB, err := ld.Load()
	if err != nil {
		fmt.Println(err)
	}

	st := storage.NewStorageProductMap(DB.DB, DB.LastId)
	vl := storage.NewStorageProductValidate(storage.ConfigStorageProductValidate{St: st, RegexExpiration: ""})
	ct := handlers.NewHandlerProducts(vl)

	rt := chi.NewRouter()

	rt.Route("/products", func(r chi.Router) {
		r.Get("/ping", ct.Ping)
		r.Post("/", ct.Create)
		r.Get("/", ct.Get)
		r.Get("/{id}", ct.GetById)
		r.Get("/search", ct.Search)
	})

	fmt.Println("Listening on port :8080")

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

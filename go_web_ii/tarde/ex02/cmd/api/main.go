package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielobarbieri/api-meli/internal/controller"
	"github.com/gabrielobarbieri/api-meli/internal/routes"
	"github.com/gabrielobarbieri/api-meli/internal/service"
	"github.com/gabrielobarbieri/api-meli/internal/storage"
	"github.com/gabrielobarbieri/api-meli/internal/storage/loader"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("%w: No .env file found, loading system environment variables", err)
		return
	}

	ld := loader.NewLoaderJSON("../../docs/db/json/products.json")
	DB, err := ld.Load()
	if err != nil {
		fmt.Println(err)
	}

	st := storage.NewStorageProductMap(DB.DB, DB.LastId)
	vl := service.NewStorageProductValidate(service.ConfigStorageProductValidate{St: st, RegexExpiration: ""})
	ct := controller.NewHandlerProducts(vl)

	rt := chi.NewRouter()

	routes.RegisterProductRoutes(rt, ct)

	fmt.Println("Listening on port :8080")

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

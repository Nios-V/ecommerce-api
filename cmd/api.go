package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Nios-V/ecommerce/api/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	productService := products.NewService()
	productsHandler := products.NewHandler(productService)
	r.Route("/products", func(r chi.Router) {
		r.Get("/", productsHandler.GetAllProducts)
	})

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.address,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server started on %s", app.config.address)

	return srv.ListenAndServe()
}

type application struct {
	config config
}

type config struct {
	address string
	db      dbConfig
}

type dbConfig struct {
	dsn string
}

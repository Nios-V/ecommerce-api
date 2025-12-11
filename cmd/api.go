package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Nios-V/ecommerce/api/internal/config"
	"github.com/Nios-V/ecommerce/api/internal/database"
	"github.com/Nios-V/ecommerce/api/internal/products"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) mount() http.Handler {

	database.Connect()
	database.Migrate(&products.Product{})

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// Products routes
	productRepo := products.NewRepository(database.DB)
	productService := products.NewService(productRepo)
	productsHandler := products.NewHandler(productService)
	r.Route("/products", func(r chi.Router) {
		r.Get("/", productsHandler.GetAllProducts)
	})

	return r
}

func (app *application) run(h http.Handler) error {
	listenAddr := ":" + app.config.ServerPort

	srv := &http.Server{
		Addr:         listenAddr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server started on %s", app.config.ServerPort)
	return srv.ListenAndServe()
}

type application struct {
	config *config.Config
}

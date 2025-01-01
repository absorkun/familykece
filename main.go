package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// load environment based on .env
	// godotenv.Load()

	// Initialize database
	initdb()

	// Configure router using chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Grouping API products
	r.Route("/api/products", func(r chi.Router) {
		r.Get("/", fetchProducts)
		r.Post("/", addProduct)
		r.Get("/{id}", showProduct)
		r.Put("/{id}", updateProduct)
		r.Delete("/{id}", deleteProduct)
	})

	// Grouping API products
	r.Route("/api/categories", func(r chi.Router) {
		r.Get("/", fetchCategories)
		r.Post("/", addCategory)
		r.Get("/{id}", showCategory)
		r.Put("/{id}", updateCategory)
		r.Delete("/{id}", deleteCategory)
	})

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	url := fmt.Sprintf("%s:%s", host, port)
	if err := http.ListenAndServe(url, r); err != nil {
		log.Fatal(err.Error())
	}
}

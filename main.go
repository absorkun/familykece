package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// Initialize database
	initdb()

	// Configure router using chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	r.Get("/", index)

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
	fmt.Printf("Server Listen at http://%s:%s or http://localhost:%s \n", host, port, port)
	if err := http.ListenAndServe(url, r); err != nil {
		log.Fatal(err.Error())
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html"))
	data := map[string]interface{}{
		"name": "absor",
	}
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func fetchProducts(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var products []Product

	if err := db.Preload("Category").Find(&products).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"data": products,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func addProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newProduct Product

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.Unmarshal(body, &newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	// newProduct.ID = r.FormValue("id")
	// newProduct.Name = r.FormValue("name")
	// priceStr := r.FormValue("price")
	// price, err := strconv.ParseFloat(priceStr, 32)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// newProduct.Price = float32(price)
	// newProduct.Icon = r.FormValue("icon")
	// newProduct.CategoryID = r.FormValue("category_id")
 //
	// if err := db.Create(&newProduct).Error; err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

}

func showProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var id string = r.PathValue("id")

	var product Product

	if err := db.Where("id = ?", id).First(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"data": product,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func updateProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var id string = r.PathValue("id")

	var newProduct Product

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.Unmarshal(body, &newProduct); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := db.Where("id = ?", id).Updates(&newProduct).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var id string = r.PathValue("id")

	var product Product

	if err := db.Where("id = ?", id).Delete(&product).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func fetchCategories(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var categories []Category

	if err := db.Find(&categories).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := map[string][]Category{
		"data": categories,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func addCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var newCategory Category

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.Unmarshal(body, &newCategory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	if err := db.Create(&newCategory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func showCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var id string = r.PathValue("id")

	var category Category

	if err := db.Where("id = ?", id).First(&category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := map[string]interface{}{
		"data": category,
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func updateCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var id string = r.PathValue("id")

	var newCategory Category

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := json.Unmarshal(body, &newCategory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if err := db.Where("id = ?", id).Updates(&newCategory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func deleteCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var id string = r.PathValue("id")

	var category Category

	if err := db.Where("id = ?", id).Delete(&category).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

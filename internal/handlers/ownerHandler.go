package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zephyr-api-mod/internal/models"
	"zephyr-api-mod/internal/service"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		http.Error(w, "specify a name for category", http.StatusBadRequest)
		return
	}
	parentId, err := strconv.Atoi(r.PathValue("parentId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.CreateCategory(name, parentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Successfully added a new category"))
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	newName := r.PathValue("newName")
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.UpdateCategory(id, newName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Successfully updated the category"))
}

func RemoveCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.RemoveCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Successfully removed the category"))
}

func GetCategories(w http.ResponseWriter, r *http.Request) {
	parentId, err := strconv.Atoi(r.PathValue("parentId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	categories, err := service.GetCategories(parentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(categories)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	product := &models.Product{}
	err := json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = service.CreateProduct(product.Name, product.InStock, product.Unit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Product added"))
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product := &models.Product{}
	err = json.NewDecoder(r.Body).Decode(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = service.UpdateProduct(id, product.Name, product.InStock, product.Unit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Product added"))
}

func RemoveProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.RemoveProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Product removed"))
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := service.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(products)
}

func AddFood(w http.ResponseWriter, r *http.Request) {
	food := &models.Food{}
	err := json.NewDecoder(r.Body).Decode(food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = service.CreateFood(food.Name, food.Price, food.Category, food.Max_quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Food added"))
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	food := &models.Food{}
	err = json.NewDecoder(r.Body).Decode(food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = service.UpdateFood(id, food.Name, food.Price, food.Max_quantity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Food changed"))
}

func RemoveFood(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = service.RemoveFood(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Food removed"))
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	food, err := service.GetFood()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(food)
}

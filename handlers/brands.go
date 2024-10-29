package handlers

import (
	"encoding/json"
	"net/http"
	"nexu-api/models"
	"nexu-api/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBrandsHandler(w http.ResponseWriter, r *http.Request) {
	brands, err := repositories.ModelRepo.GetAllBrands()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(brands)
}

func CreateBrandHandler(w http.ResponseWriter, r *http.Request) {
	var brand models.Brand
	if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	createdBrand, err := repositories.ModelRepo.CreateBrand(brand)
	if err != nil {
		if createdBrand.ID == 0 {
			http.Error(w, "Brand already exists", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBrand)
}

func GetBrandModelsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brandID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid brand ID", http.StatusBadRequest)
		return
	}
	models, err := repositories.ModelRepo.GetModelsByBrandID(brandID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models)
}

func CreateBrandModelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brandID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid brand ID", http.StatusBadRequest)
		return
	}
	var model models.ModelDB
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if model.AveragePrice < 100000 && model.AveragePrice != 0 {
		http.Error(w, "Average price must be greater than 100000", http.StatusBadRequest)
		return
	}

	createdModel, err := repositories.ModelRepo.CreateModelForBrand(brandID, model)
	if err != nil {
		if createdModel.ID == 0 {
			http.Error(w, " Brand not found", http.StatusBadRequest)
			return
		}
		if createdModel.ID == -1 {
			http.Error(w, "Model already exists", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdModel)
}

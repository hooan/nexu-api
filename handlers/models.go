package handlers

import (
	"encoding/json"
	"net/http"
	"nexu-api/models"
	"nexu-api/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

func GetModelsHandler(w http.ResponseWriter, r *http.Request) {

	greater, err := strconv.ParseFloat(r.URL.Query().Get("greater"), 32)
	if err != nil {
		http.Error(w, "Invalid 'greater' parameter", http.StatusBadRequest)
		return
	}
	lower, err := strconv.ParseFloat(r.URL.Query().Get("lower"), 32)
	if err != nil {
		http.Error(w, "Invalid 'lower' parameter", http.StatusBadRequest)
		return
	}
	filter := models.Filter{Greater: float32(greater), Lower: float32(lower)}

	models, err := repositories.ModelRepo.GetAllModels(filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models)
}

func UpdateModelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	modelID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid model ID", http.StatusBadRequest)
		return
	}

	var model models.Model
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	model.ID = modelID

	if model.AveragePrice < 100000 {
		http.Error(w, "Average price must be greater than 100000", http.StatusBadRequest)
		return
	}

	if err := repositories.ModelRepo.UpdateModel(model); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

package Infrastructure

import (
	foodApp "FoodInHouse/Food/Application"
	foodService "FoodInHouse/Food/Domain"
	"encoding/json"
	"net/http"
)

func GetFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	foods, err := foodApp.Read()

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(foods)
}

func GetFoodById(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	w.Header().Set("Content-Type", "application/json")
	foods, err := foodApp.ReadById(ID)

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(foods)
}

func CreateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m foodService.Food

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	err = foodService.Create(m)

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")

	var m foodService.Food

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	err = foodService.Update(m, ID)

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

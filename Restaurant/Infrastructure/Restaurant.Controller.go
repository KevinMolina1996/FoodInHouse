package Infrastructure

import (
	restaurantApp "FoodInHouse/Restaurant/Application"
	restaurantService "FoodInHouse/Restaurant/Domain"
	"encoding/json"
	"net/http"
)

func GetRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurants, err := restaurantApp.Read()

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(restaurants)
}

func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m restaurantService.Restaurant

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	err = restaurantApp.Create(m)

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	w.Header().Set("Content-Type", "application/json")

	var m restaurantService.Restaurant

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	err = restaurantApp.Update(m, ID)

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

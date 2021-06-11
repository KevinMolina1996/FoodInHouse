package Infrastructure

import (
	orderApp "FoodInHouse/Order/Application"
	orderService "FoodInHouse/Order/Domain"
	"encoding/json"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m orderService.Order

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	err = orderApp.Create(m)

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(m)
}

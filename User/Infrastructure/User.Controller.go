package Infrastructure

import (
	userApp "FoodInHouse/User/Application"
	userService "FoodInHouse/User/Domain"
	"encoding/json"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	restaurants, err := userApp.Read()

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(restaurants)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var m userService.User

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	err = userApp.Create(m)

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ID := r.URL.Query().Get("id")

	var m userService.User

	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		http.Error(w, "", 400)
		return
	}

	err = userApp.Update(m, ID)

	if err != nil {
		http.Error(w, "Error en la base de datos, desc: "+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

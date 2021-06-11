package Handlers

import (
	"log"
	"net/http"

	foodRepository "FoodInHouse/Food/Infrastructure"
	orderRepository "FoodInHouse/Order/Infrastructure"
	restaurantRepository "FoodInHouse/Restaurant/Infrastructure"
	userRepository "FoodInHouse/User/Infrastructure"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Drivers() {
	r := mux.NewRouter()
	r.HandleFunc("/api/getUsers", userRepository.GetUsers).Methods("GET")
	r.HandleFunc("/api/createUser", userRepository.CreateUser).Methods("PUT")
	r.HandleFunc("/api/updateUser", userRepository.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/getRestaurants", restaurantRepository.GetRestaurants).Methods("GET")
	r.HandleFunc("/api/createRestaurant", restaurantRepository.CreateRestaurant).Methods("PUT")
	r.HandleFunc("/api/updateRestaurant", restaurantRepository.UpdateRestaurant).Methods("PUT")
	r.HandleFunc("/api/getFoodById", foodRepository.GetFoodById).Methods("GET")
	r.HandleFunc("/api/createFood", foodRepository.CreateFood).Methods("PUT")
	r.HandleFunc("/api/updateFood", foodRepository.UpdateFood).Methods("PUT")
	r.HandleFunc("/api/createOrder", orderRepository.CreateOrder).Methods("POST")
	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

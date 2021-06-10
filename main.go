package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome to FoodInHouse API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	fmt.Println("Server run on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

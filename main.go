package main

import (
	"fullcycle_cep_weather/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting server...")
	defer log.Println("Server finished...")

	r := mux.NewRouter()
	r.HandleFunc("/weather/{cep}", handler.GetWeatherHandler).Methods(http.MethodGet)

	log.Println("listening on port 8080")
	http.ListenAndServe(":8080", r)
}

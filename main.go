package main

import (
	"log"
	"net/http"
	"os"

	"github.com/_url-Short-App/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Home).Methods("GET")
	r.HandleFunc("/GenerateShortUrl", handlers.GenerateShortUrl).Methods("POST")
	r.HandleFunc("/RestoreUrl", handlers.RestoreUrl).Methods("GET")
	r.HandleFunc("/DeleteUrl", handlers.DeleteUrl).Methods("GET")

	r.HandleFunc("/Get", handlers.GetbyId).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Listening...")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

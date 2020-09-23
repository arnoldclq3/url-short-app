package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/_url-Short-App/business"
	"github.com/_url-Short-App/services"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", Home).Methods("GET")
	r.HandleFunc("/GenerateShortUrl", GenerateShortUrl).Methods("GET")
	r.HandleFunc("/Get", GetbyId).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Listening...")
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	welcome := "Hola mundo! yo soy Arnold"
	respondWithJSON(w, http.StatusOK, welcome)
}

func GenerateShortUrl(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "URL no valida")
	}

	var url string
	json.Unmarshal(reqBody, &url)

	myShortener := services.NewShortenerBase26()
	mydb := services.NewMongoService()
	// mydb := services.NewMockDataBase()
	urlBiz := business.NewUrlBusiness(*myShortener, mydb)

	shorUrl := urlBiz.GenerateShortURL(url)

	result := "La url generada es: " + string(shorUrl)
	respondWithJSON(w, http.StatusOK, result)
}

func GetbyId(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "URL no valida")
	}

	var id int
	json.Unmarshal(reqBody, &id)

	myShortener := services.NewShortenerBase26()
	mydb := services.NewMongoService()
	// mydb := services.NewMockDataBase()
	urlBiz := business.NewUrlBusiness(*myShortener, mydb)

	result, err := urlBiz.GetbyId(id)
	_ = err

	respondWithJSON(w, http.StatusOK, result)
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

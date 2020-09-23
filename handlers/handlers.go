package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/_url-Short-App/business"
	"github.com/_url-Short-App/services"
)

type UrlDel struct {
	Url      string
	ShortUrl string
}

func Home(w http.ResponseWriter, r *http.Request) {
	nombre := "Arnold"

  // welcome := "Hola mundo! yo soy " + nombre + " Y esta es mi 1ra app en GO"

	welcome := struct {
		MiNombre   string
		Mensaje    string
		VersionApp string
	}{
		"Hola mundo! yo soy " + nombre,
		"Y esta es mi 1ra app en GO",
		"v1.7",
	}

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
	urlBiz := business.NewUrlBusiness(*myShortener, mydb)

	shortUrl := urlBiz.GenerateShortURL(url)

	result := struct {
		UrlGen string
	}{
		shortUrl,
	}

	respondWithJSON(w, http.StatusOK, result)
}

func RestoreUrl(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "URL no valida")
	}

	var url string
	json.Unmarshal(reqBody, &url)

	myShortener := services.NewShortenerBase26()
	mydb := services.NewMongoService()
	urlBiz := business.NewUrlBusiness(*myShortener, mydb)

	originUrl := urlBiz.RestoreOriginalURL(url)

	result := struct {
		Url string
	}{
		originUrl,
	}

	respondWithJSON(w, http.StatusOK, result)
}

func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "URL no valida")
	}

	var urlDel UrlDel
	err = json.Unmarshal(reqBody, &urlDel)
	if err != nil {
		fmt.Fprintf(w, "Tipo no Valido")
	}

	myShortener := services.NewShortenerBase26()
	mydb := services.NewMongoService()
	urlBiz := business.NewUrlBusiness(*myShortener, mydb)
	result := "Eliminado"
	if urlDel.Url != "" {
		err = urlBiz.DeleteByUrl(urlDel.Url)
	} else if urlDel.ShortUrl != "" {
		err = urlBiz.DeleteByShortUrl(urlDel.ShortUrl)
	} else {
		result = "url no existe"
	}

	if err != nil {
		result = "Error al eliminar"
	}

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

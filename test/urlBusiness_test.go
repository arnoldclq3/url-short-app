package test

import (
	"testing"

	"github.com/_url-Short-App/business"
	"github.com/_url-Short-App/services"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURL(t *testing.T) {
	myShortener := services.NewShortenerBase26()
	mydb := services.NewMockDataBase()
	myBusiness := business.NewUrlBusiness(*myShortener, mydb)

	result := myBusiness.GenerateShortURL("wwwwwwwwwwwwwwwwwwwwwwwwww")

	assert.Equal(t, "B", result, "fallo test GenerateShortURL")
}

func TestRestoreOriginalURL(t *testing.T) {
	myShortener := services.NewShortenerBase26()
	mydb := services.NewMockDataBase()
	myBusiness := business.NewUrlBusiness(*myShortener, mydb)

	urlGen := myBusiness.GenerateShortURL("wwwwwwwwwwwwwwwwwwwwwwwwww")
	result := myBusiness.RestoreOriginalURL(urlGen)

	assert.Equal(t, "wwwwwwwwwwwwwwwwwwwwwwwwww", result, "fallo test GenerateShortURL")
}

func TestRestoreInexistentURL(t *testing.T) {
	myShortener := services.NewShortenerBase26()
	mydb := services.NewMockDataBase()
	myBusiness := business.NewUrlBusiness(*myShortener, mydb)

	result := myBusiness.RestoreOriginalURL("AXVCVX")

	assert.Equal(t, "", result, "fallo test TestRestoreInexistentURL")
}

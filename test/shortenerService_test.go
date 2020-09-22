package test

import (
	"testing"

	"github.com/_url-Short-App/services"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortString(t *testing.T) {
	// myService := new(services.ShortenerService)
	myService := services.NewShortenerService()

	result := myService.GenerateShortString(125)

	assert.Equal(t, "EV", result, "fallo test GenerateShortString")
}

func TestRestoreSeedFromString(t *testing.T) {
	// myService := new(services.ShortenerService)
	myService := services.NewShortenerService()

	result := myService.RestoreSeedFromString("EV")

	assert.Equal(t, 125, result, "fallo test RestoreSeedFromString")
}

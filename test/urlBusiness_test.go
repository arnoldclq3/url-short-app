package test

import (
	"testing"

	"github.com/_url-Short-App/business"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURL(t *testing.T) {

	myBusiness := business.NewUrlBusiness()

	result := myBusiness.GenerateShortURL("wwwwwwwwwwwwwwwwwwwwwwwwww")

	assert.Equal(t, "A", result, "fallo test GenerateShortURL")
}

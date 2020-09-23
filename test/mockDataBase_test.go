package test

import (
	"testing"

	"github.com/_url-Short-App/entities"
	"github.com/_url-Short-App/services"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	mydb := services.NewMockDataBase()

	item := entities.Url{Id: -1, Text: "wwwwwwwwwwwwwwwwwwww"}
	mydb.Add(item)
	urls, err := mydb.GetAll()
	_ = err

	assert.Equal(t, 1, len(urls), "fallo test TestAdd")
}

func TestFindById(t *testing.T) {
	mydb := services.NewMockDataBase()

	item := entities.Url{Id: 1, Text: "wwwwwwwwwwwwwwwwwwww"}
	item2 := entities.Url{Id: 2, Text: "otraurlsuperlarga"}
	mydb.Add(item)
	mydb.Add(item2)
	url, err := mydb.FindById(2)
	_ = err

	assert.Equal(t, "otraurlsuperlarga", url.Text, "fallo test FindById")
}

func TestFindByText(t *testing.T) {
	mydb := services.NewMockDataBase()

	item := entities.Url{Id: -1, Text: "wwwwwwwwwwwwwwwwwwww"}
	item2 := entities.Url{Id: -1, Text: "otraurlsuperlarga"}
	item3 := entities.Url{Id: -1, Text: "4544444444444444443444"}
	mydb.Add(item)
	mydb.Add(item2)
	mydb.Add(item3)
	url, err := mydb.Find(item2)
	_ = err

	assert.Equal(t, "otraurlsuperlarga", url.Text, "fallo test Find")
}

func TestDelete(t *testing.T) {
	mydb := services.NewMockDataBase()

	item := entities.Url{Id: 1, Text: "wwwwwwwwwwwwwwwwwwww"}
	item2 := entities.Url{Id: 2, Text: "otraurlsuperlarga"}
	item3 := entities.Url{Id: 3, Text: "4544444444444444443444"}
	mydb.Add(item)
	mydb.Add(item2)
	mydb.Add(item3)
	err := mydb.Delete(2)
	_ = err
	urls, err := mydb.GetAll()
	_ = err

	assert.Equal(t, 2, len(urls), "fallo test Delete")
}

func TestFindNotExist(t *testing.T) {
	mydb := services.NewMockDataBase()

	item := entities.Url{Id: -1, Text: "wwwwwwwwwwwwwwwwwwww"}
	item2 := entities.Url{Id: -1, Text: "otraurlsuperlarga"}
	item3 := entities.Url{Id: -1, Text: "4544444444444444443444"}
	mydb.Add(item)
	mydb.Add(item2)
	mydb.Add(item3)
	url, err := mydb.FindById(454)
	_ = url

	assert.NotNil(t, err, "fallo test Delete")
}

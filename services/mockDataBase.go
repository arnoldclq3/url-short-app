package services

import (
	"errors"

	"github.com/_url-Short-App/entities"
)

type MockDataBase struct {
	urls []entities.Url
}

func NewMockDataBase() *MockDataBase {
	return &MockDataBase{urls: make([]entities.Url, 0)}
}

func (db *MockDataBase) FindById(id int) (entities.Url, error) {
	var result entities.Url
	for i := range db.urls {
		if db.urls[i].Id == id {
			return db.urls[i], nil
		}
	}
	return result, errors.New("Element not exist")
}

func (db *MockDataBase) Find(anyUrl entities.Url) (entities.Url, error) {
	var result entities.Url
	for i := range db.urls {
		if db.urls[i].Text == anyUrl.Text {
			return db.urls[i], nil
		}
	}
	return result, errors.New("Element not exist")
}

func (db *MockDataBase) Add(anyUrl entities.Url) error {
	anyUrl.Id = len(db.urls)
	db.urls = append(db.urls, anyUrl)
	return nil
}

func (db *MockDataBase) Update(id int, anyUrl entities.Url) error {
	return nil
}

func (db *MockDataBase) Delete(id int) error {
	for i := range db.urls {
		if db.urls[i].Id == id {
			db.urls = append(db.urls[:i], db.urls[i+1:]...)
			break
		}
	}
	return nil
}

func (db *MockDataBase) GetAll() ([]entities.Url, error) {
	return db.urls, nil
}

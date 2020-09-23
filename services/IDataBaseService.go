package services

import "github.com/_url-Short-App/entities"

type IDataBaseService interface {
	FindById(int) (entities.Url, error)
	Find(entities.Url) (entities.Url, error)
	Add(entities.Url) error
	Update(int, entities.Url) error
	Delete(int) error
	GetAll() ([]entities.Url, error)
	FindLast() (entities.Url, error)
}

package business

import (
	"github.com/_url-Short-App/entities"
	"github.com/_url-Short-App/services"
)

type UrlBusiness struct {
	shortener services.IShortenerService
	db        services.IDataBaseService
}

func NewUrlBusiness(shortener services.IShortenerService, db services.IDataBaseService) *UrlBusiness {
	u := new(UrlBusiness)
	u.shortener = shortener
	u.db = db
	return u
}

func (u UrlBusiness) GenerateShortURL(longUrl string) string {
	last, err := u.db.FindLast()
	if err != nil {
		return ""
	}

	idGen := last.Id + 1
	url := entities.Url{Id: idGen, Text: longUrl}

	err = u.db.Add(url)
	if err != nil {
		return ""
	}

	shortUrlGen := u.shortener.GenerateShortString(idGen)
	return shortUrlGen
}

func (u UrlBusiness) RestoreOriginalURL(shortUrl string) string {
	idRec := u.shortener.RestoreSeedFromString(shortUrl)

	urlRec, err := u.db.FindById(idRec)
	if err != nil {
		return ""
	}

	urlOriginal := urlRec.Text
	return urlOriginal
}

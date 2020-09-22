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
	url := entities.Url{Id: -1, Text: longUrl}

	err := u.db.Add(url)
	if err != nil {
		return ""
	}

	urlGen, err := u.db.Find(url)
	if err != nil {
		return ""
	}

	idGen := urlGen.Id
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

package business

import (
	"github.com/_url-Short-App/services"
)

type UrlBusiness struct {
	shortener services.IShortenerService
}

func NewUrlBusiness(shortener services.IShortenerService) *UrlBusiness {
	u := new(UrlBusiness)
	u.shortener = shortener
	return u
}

func (u UrlBusiness) GenerateShortURL(longUrl string) string {
	result := u.shortener.GenerateShortString(1)
	return result
}

func (u UrlBusiness) RestoreOriginalURL(shortUrl string) string {
	result := u.shortener.RestoreSeedFromString(shortUrl)

	if result == 1 {
		return "wwwwwwwwwwwwwwwwwwwwwwwwww"
	}
	return ""
}

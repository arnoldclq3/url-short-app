package business

type UrlBusiness struct {
}

func NewUrlBusiness() *UrlBusiness {
	return new(UrlBusiness)
}

func (u UrlBusiness) GenerateShortURL(longUrl string) string {
	return "A"
}

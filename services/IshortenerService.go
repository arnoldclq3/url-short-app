package services

type IShortenerService interface {
	GenerateShortString(int) string
	RestoreSeedFromString(string) int
}

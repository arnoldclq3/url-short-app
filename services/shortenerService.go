package services

import "math"

type ShortenerService struct {
	alphabet string
	base     int
	dicc     map[rune]int
}

func (s ShortenerService) GenerateShortString(seed int) string {
	s.dicc = make(map[rune]int)
	s.alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s.base = len(s.alphabet)
	for pos, char := range s.alphabet {
		s.dicc[char] = pos
	}

	if seed < s.base {
		return string(s.alphabet[0])
	}

	str := ""
	i := seed

	for i > 0 {
		str = string(s.alphabet[i%s.base]) + str
		i /= s.base
	}

	return str
	//return "EV"
}

func (s ShortenerService) RestoreSeedFromString(seed string) int {
	s.dicc = make(map[rune]int)
	s.alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s.base = len(s.alphabet)
	for pos, char := range s.alphabet {
		s.dicc[char] = pos
	}

	result := 0
	for pos, char := range seed {
		result += s.dicc[char] * pow(s.base, len(seed)-1-pos)
	}
	return result
	// return 125
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

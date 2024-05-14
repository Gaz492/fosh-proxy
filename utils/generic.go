package utils

import (
	"os"
	"unicode"
	"unicode/utf8"
)

func ToLowerFirstChar(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func FarenheitToCelsius(tempF float64) float64 {
	//tempF, err := strconv.ParseFloat(tempFStr, 64)
	//if err != nil {
	//	return ""
	//}
	tempC := (tempF - 32) * 5 / 9
	//tempCStr := fmt.Sprintf("%f", tempC)
	return tempC
}

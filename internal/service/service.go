package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(data string) (string, error) {
	data = strings.TrimSpace(data)
	if data == "" {
		return "", errors.New("пустые данные")
	}

	if isMorse(data) {
		return morse.ToText(data), nil
	}
	return morse.ToMorse(data), nil
}
func isMorse(s string) bool {
	for _, r := range s {
		switch r {
		case '.', '-', ' ':
		default:
			return false
		}
	}
	return true
}

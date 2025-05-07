package service

import (
	"errors"
	"strings"

	"github.com/Kras0Tanya/morse-converter-by-KisaSoft/pkg/morse"
)

// это самое злодремучее финальное задание; страшно представить, что будет дальше... :`(

func looksLikeMorse(data string) bool {

	for _, r := range data {
		if r != '.' && r != '-' && r != ' ' {
			return false
		}
	}
	return true
}

func DetectAndConvert(data string) (string, error) {
	trimmed := strings.TrimSpace(data)

	if trimmed == "" {
		return "", errors.New("input string is empty")
	}

	norm := strings.ReplaceAll(trimmed, "\n", "")
	normData := strings.ReplaceAll(norm, "\r", "")

	if looksLikeMorse(normData) {
		return morse.ToText(normData), nil
	}
	return morse.ToMorse(normData), nil
}

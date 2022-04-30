package utils

import (
	"math/rand"
	"strings"
)

func RandString(length int, ascii, digits, punctuation bool) string {
	var charBuilder strings.Builder
	var result strings.Builder

	if ascii {
		charBuilder.WriteString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if digits {
		charBuilder.WriteString("0123456789")
	}
	if punctuation {
		charBuilder.WriteString("!%()*+,-.:;<=>?@[]^_{|}~")
	}
	chars := charBuilder.String()

	for i := 0; i < length; i++ {
		random := rand.Intn(len(chars))
		randomChar := chars[random]
		result.WriteString(string(randomChar))
	}

	return result.String()
}

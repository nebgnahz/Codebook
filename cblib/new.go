package cblib

import (
	"crypto/rand"
)

const (
	Numbers string = "0123456789"
	LETTERS string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters string = "abcdefghijklmnopqrstuvwxyz"
)

func NewPasscode(n int) string {
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = LETTERS[b%byte(len(LETTERS))]
	}
	return string(bytes)
}

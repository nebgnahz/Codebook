package cblib

import (
	"crypto/rand"
)

const (
	Numbers string = "0123456789"
	LETTERS string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letters string = "abcdefghijklmnopqrstuvwxyz"
	Symbols string = ",.!@#$%^&*()"
)

func NewPasscodeEasy(length int) string {
	return randomStringsFromStrings(Numbers+letters, length)
}

func NewPasscodeNormal(length int) string {
	return randomStringsFromStrings(Numbers+LETTERS+letters, length)
}

func NewPasscodeHard(length int) string {
	return randomStringsFromStrings(Numbers+LETTERS+letters+Symbols, length)
}

func randomStringsFromStrings(in string, length int) string {
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = in[b%byte(len(in))]
	}
	return string(bytes)
}

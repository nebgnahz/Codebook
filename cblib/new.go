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

func NewPasscodeEasy(length int) []byte {
	return randomStringsFromStrings(Numbers+letters, length)
}

func NewPasscodeNormal(length int) []byte {
	return randomStringsFromStrings(Numbers+LETTERS+letters, length)
}

func NewPasscodeHard(length int) []byte {
	return randomStringsFromStrings(Numbers+LETTERS+letters+Symbols, length)
}

func randomStringsFromStrings(in string, length int) []byte {
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = in[b%byte(len(in))]
	}
	return bytes
}

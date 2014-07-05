package cblib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func KeyNormalize(key []byte) []byte {
	// KeyNormalize will
	//   1, pad the key if length too small
	//   2, truncate the key if length too large
	if len(key) > aes.BlockSize {
		return key[0:aes.BlockSize]
	}
	new_key := make([]byte, aes.BlockSize, aes.BlockSize)
	copy(new_key, key)
	return new_key
}

func EncodeBase64(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func DecodeBase64(b []byte) []byte {
	data, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	b := EncodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return ciphertext
}

func Decrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return DecodeBase64(text)
}

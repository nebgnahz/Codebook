package cblib

import (
	"bytes"
	"testing"
)

func TestKeyTruncating(t *testing.T) {
	key := []byte("A secret and very long key")
	new_key := KeyNormalize(key)
	expected_key := []byte("A secret and ver")

	if !bytes.Equal(new_key, expected_key) {
		t.Error(
			"Key", key,
			"expected", expected_key,
			"got", new_key,
		)
	}
}

func TestKeyPadding(t *testing.T) {
	key := []byte("A secret key")
	new_key := KeyNormalize(key)
	expected_key := []byte("A secret key\x00\x00\x00\x00")

	if !bytes.Equal(new_key, expected_key) {
		t.Error(
			"Key", key,
			"expected", expected_key,
			"got", new_key,
		)
	}
}

func TestEncryptDecrypt(t *testing.T) {
	key := []byte("A secret key")
	plaintext := []byte("After all is said and done, more is said than done.")

	new_key := KeyNormalize(key)
	ciphertext := Encrypt(new_key, plaintext)
	result := Decrypt(new_key, ciphertext)
	if !bytes.Equal(result, plaintext) {
		t.Error("Encryption Failed")
	}
}

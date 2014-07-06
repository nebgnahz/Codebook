package cblib

import (
	"bytes"
	"fmt"
	"errors"
)

func (c *Codebook) Get(website []byte) ([]byte, error) {
	// for now, let's just iterate through the codes
	for _, kv := range c.codes {
		if bytes.Equal(kv.key, website) {
			pwd := Decrypt(c.masterkey, kv.value)
			return pwd, nil
		}
	}
	return nil, errors.New("website not found")
}

func (c *Codebook) PrintPlain() {
	// for now, let's just iterate through the codes
	for _, kv := range c.codes {
		decrypted := Decrypt(c.masterkey, kv.value)
		fmt.Println(
			"key:",
			string(kv.key),
			" | value:",
			string(decrypted))
	}
}

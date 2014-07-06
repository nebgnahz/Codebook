package cblib

import (
	"fmt"
)

func (c *Codebook) Add(website, password []byte) {
	fmt.Println(string(website), string(password))
	encrypted_password := Encrypt(c.masterkey, password)
	kv := KV{key: website, value: encrypted_password}
	c.codes = append(c.codes, kv)
}

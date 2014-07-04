package cblib

import (
	"fmt"
)

func (c *Codebook) Add(website, password string) {
	fmt.Println(website, password)

	encrypted_password := string(Encrypt([]byte(c.masterkey), []byte(password)))
	kv := KV{key: website, value: encrypted_password}
	c.codes = append(c.codes, kv)
}

package cblib

import (
	"bytes"
	"fmt"
)

func (c *Codebook) Get(website []byte) bool {
	// for now, let's just iterate through the codes
	for _, kv := range c.codes {
		if bytes.Equal(kv.key, website) {
			return true
		}
	}
	return false
}

func (c *Codebook) PrintPlain() {
	// for now, let's just iterate through the codes
	for _, kv := range c.codes {
		fmt.Println("key:", kv.key, " | value:", kv.value)
	}
}

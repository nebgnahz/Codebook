package cblib

import (
	"fmt"
)

func (c *Codebook) Get(website string) bool {
	// for now, let's just iterate through the codes
	for _, kv := range c.codes {
		if kv.key == website {
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

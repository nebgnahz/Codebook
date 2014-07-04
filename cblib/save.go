package cblib

import (
	"os"
)

func (c *Codebook) Save() {
	f, err := os.Create(CodebookFile)
	check(err)
	defer f.Close()

	for _, kv := range c.codes {
		// format being: key:value
		_, err := f.Write([]byte(kv.key))
		check(err)
		_, err = f.Write([]byte(":"))
		check(err)
		_, err = f.Write([]byte(kv.value))
		check(err)
	}
}

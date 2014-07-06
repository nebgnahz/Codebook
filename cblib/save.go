package cblib

import (
	"os"
)

func (c *Codebook) Save() {
	f, err := os.Create(c.bookfile)
	check(err)
	defer f.Close()

	for _, kv := range c.codes {
		// format being: key:value
		_, err := f.Write(kv.key)
		check(err)
		_, err = f.Write([]byte(":"))
		check(err)
		_, err = f.Write(EncodeBase64(kv.value))
		check(err)
		f.Write([]byte("\n"))
		check(err)
	}
}

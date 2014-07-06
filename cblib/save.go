package cblib

import (
	"crypto/sha1"
	"os"
)

func (c *Codebook) Save() {
	f, err := os.Create(c.bookfile)
	check(err)
	defer f.Close()

	// Save the first line as masterkey digest
	sha := sha1.Sum(c.masterkey)
	_, err = f.Write(EncodeBase64(sha[:len(sha)]))
	check(err)
	_, err = f.Write([]byte("\n"))
	check(err)
	
	for _, kv := range c.codes {
		// format being: key:value
		_, err := f.Write(kv.key)
		check(err)
		_, err = f.Write([]byte(":"))
		check(err)
		_, err = f.Write(EncodeBase64(kv.value))
		check(err)
		_, err = f.Write([]byte("\n"))
		check(err)
	}
}

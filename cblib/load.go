package cblib

import (
	"bytes"
	"fmt"
	"os"
)

// Load will read the file and parse it
func (c *Codebook) Load() {
	c.codes = make([]KV, 0)
	// Make sure we can read the file
	if _, err := os.Stat(c.bookfile); err == nil {
		// load the file
		f, _ := os.Open(c.bookfile)
		defer f.Close()

		data := make([]byte, 1000)
		if count, err := f.Read(data); err == nil {
			fmt.Printf("read %d bytes: %q\n", count, data[:count])
			for _, entry := range bytes.Split(data[:count], []byte("\n")) {
				if bytes.Index(entry, []byte(":")) == -1 {
					break
				}
				pair := bytes.SplitN(entry, []byte(":"), 2)
				kv := KV{
					key:   pair[0],
					value: DecodeBase64(pair[1]),
				}
				c.codes = append(c.codes, kv)
			}
		}
	}
}

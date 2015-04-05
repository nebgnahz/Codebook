package cblib

import (
	"bytes"
	"crypto/sha1"
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

		// TODO(benzh) Fix this, the read to data is not scalable to handle large
		// files.
		data := make([]byte, 50000)
		if count, err := f.Read(data); err == nil {
			// log.Println("read", count, "bytes:", data[:count])
			for i, entry := range bytes.Split(data[:count], []byte("\n")) {

				// first line being sha1 sum
				sha := sha1.Sum(c.masterkey)
				if i == 0 {
					if bytes.Equal(sha[:len(sha)], DecodeBase64(entry)) {
						continue
					} else {
						panic("master key doens't match")
					}
				}

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

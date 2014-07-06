package cblib

import (
	"os"
)

type KV struct {
	key   []byte
	value []byte
}

type Codebook struct {
	masterkey []byte
	codes     []KV
	bookfile  string
}

const (
	CodebookFile = ".codebook"
)

func Init(master_key string) *Codebook {
	filepath := os.Getenv("HOME") + "/" + CodebookFile
	c := &Codebook{
		masterkey: KeyNormalize([]byte(master_key)),
		bookfile:  filepath,
	}
	c.Load()
	return c
}

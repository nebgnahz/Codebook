package cblib

import ()

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
	CodebookFile = "/tmp/.codebook"
)

func Init(master_key string) *Codebook {
	c := &Codebook{
		masterkey: KeyNormalize([]byte(master_key)),
		bookfile:  CodebookFile,
	}
	c.Load()
	return c
}

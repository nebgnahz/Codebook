package cblib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type KV struct {
	key   string
	value string
}

type Codebook struct {
	masterkey string
	codes     []KV
}

const (
	CodebookFile = "/Users/benzh/.codebook"
)

func Init(master_key string) *Codebook {
	master_key = Pad(master_key)

	m := &Codebook{masterkey: master_key}
	m.codes = make([]KV, 0)
	// initialization will create a file for key-value store
	if _, err := os.Stat(CodebookFile); err == nil {
		// load the file
		f, _ := os.Open(CodebookFile)
		defer f.Close()
		r := bufio.NewReader(f)
		for err == nil {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			}
			// parse the line
			s := strings.TrimRight(string(line), "\n")
			pair := strings.SplitN(s, ":", 2)
			// prepare codebook
			kv := KV{key: pair[0], value: string(Decrypt([]byte(master_key), []byte(pair[1])))}

			fmt.Println(kv)

			m.codes = append(m.codes, kv)
		}
	}
	return m
}

package cblib

// A medium size system testing
// 1, we generate new keys and save them into temporary files
// 2, we read the temporary file and parse it to compare the value
// 3, clean up

import (
	"os/exec"
	"bytes"
	"testing"
)

const (
	TmpCodebookFile = "/tmp/codebook"
	TmpMasterKey    = "test key"
)

var expected = make([]KV, 0)

func TestLoading(t *testing.T) {
	GenerateNewFile()
	c := LoadFromFile()
	
	// Make comparison
	for i, kv := range c.codes {
		decrypted := Decrypt(c.masterkey, kv.value)
		if !bytes.Equal(kv.key, expected[i].key) ||
			!bytes.Equal(decrypted, expected[i].value) {
			t.Error(
				"expected", expected[i],
				"got", kv,
			)
		}
	}
	
	CleanUp()
}

func GenerateNewFile() {
	c := &Codebook{
		masterkey: KeyNormalize([]byte(TmpMasterKey)),
		bookfile:  TmpCodebookFile,
	}
	key1 := []byte("baidu.com")
	value1 := NewPasscodeHard(15)
	expected = append(expected, KV{key: key1, value: value1})
	c.Add(key1, value1)

	key2 := []byte("youku.com")
	value2 := NewPasscodeHard(19)
	expected = append(expected, KV{key: key2, value: value2})
	c.Add(key2, value2)

	c.Save()
}

func LoadFromFile() *Codebook {
	c := &Codebook{
		masterkey: KeyNormalize([]byte(TmpMasterKey)),
		bookfile:  TmpCodebookFile,
	}
	c.Load()
	return c
}

func CleanUp() {
	c1 := exec.Command("rm", TmpCodebookFile)
	c1.Start()
}

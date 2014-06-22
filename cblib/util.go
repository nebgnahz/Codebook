package cblib

import (
	"io"
	"os/exec"
)

func CopyToClipBoard(str string) {
	// uses Mac OS shell command pbcopy
	c1 := exec.Command("echo", str)
	c2 := exec.Command("pbcopy")

	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
}
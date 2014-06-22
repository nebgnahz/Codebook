// ok, top down approach
// this cli will take a few routes:
// 1, new, when you register a website
// 2, add, if you manually created the password and want it to be stored here
// 3, set, you may update the value
// 4, get, get the password
// 5, rm, removes the entry
// 6, ls, list what's stored
// when we create the file, we should consider encrypt it with a master key. let's do this later

// so far the following is only a hello world application that takes arguments...
package main

import (
	"codebook/cblib"
	"flag"
	"fmt"
)

var command = flag.String("help", "help", "Codebook is the tool to manage your passcode for all websites")
var website = flag.String("website", "google", "the website you need to know your secret password")

func main() {
	x := cblib.NewPasscodeHard(15)
	flag.Parse()

	fmt.Println("The password for", *website, "is", x)
}

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

var command = flag.String("command", "help", "Codebook is the tool to manage your passcode for all websites")

func main() {
	flag.Parse()
	switch *command {
	case "help":
		fmt.Println("TODO(benzh), print all available handlers")
	case "test":
		c := cblib.Init("key")
		c.Add("baidu.com", cblib.NewPasscodeHard(15))
		c.Save()
		fmt.Println(c)
	case "print":
		c := cblib.Init("key")
		c.PrintPlain()
	case "new":
		fmt.Println("Enter the website:")
		var website, y_or_n string
		_, _ = fmt.Scanf("%s", &website)
		pc := cblib.NewPasscodeHard(15)
		fmt.Println("The password for", website, "is", pc)
		fmt.Println("Accept? (y/N):")
		_, _ = fmt.Scanf("%s", &y_or_n)
		if y_or_n == "y" {
			cblib.CopyToClipBoard(pc)
		}
	}
}

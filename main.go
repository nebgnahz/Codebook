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
	"os"
)

const (
	VERSION = "0.1"
)

func PrintUsage() {
	fmt.Print(
		`usage: codebook [--version] [--help] <command> [<args>]

The most common codebook commands are:
  new      generate a random password for the new website
  add      manually add another entry for the website
  set      manually set the password for a website
  get      return the password for a specific website   
`)
}

func main() {
	if len(os.Args) < 2 {
		PrintUsage()
		return
	}

	if FlagParsing() {
		return
	}

	fmt.Println("Enter master key (recommended shorter than 16 bytes):")
	var master_key string
	_, _ = fmt.Scanf("%s", &master_key)
	c := cblib.Init(master_key)

	switch os.Args[1] {
	case "get":
		pwd := make([]byte, 0)
		if c.Get([]byte(os.Args[2]), pwd) {
			fmt.Println(pwd)
		} else {
			fmt.Println(os.Args[2], "not found")
		}
	case "new":
		fmt.Println("Enter the website:")
		var website, y_or_n string
		_, _ = fmt.Scanf("%s", &website)
		pc := cblib.NewPasscodeHard(15)
		fmt.Println("The password for", website, "is", pc)
		fmt.Println("Accept? (y/N):")
		_, _ = fmt.Scanf("%s", &y_or_n)
		if y_or_n == "y" {
			cblib.CopyToClipBoard(string(pc))
		}
	}
}

func FlagParsing() bool {
	// flag parsing
	var help = flag.Bool("help", false, "Codebook is the tool to manage your passcode for all websites.")
	var version = flag.Bool("version", false, "")

	flag.Parse()

	if *help {
		PrintUsage()
		return true
	} else if *version {
		fmt.Println(VERSION)
		return true
	}
	return false
}

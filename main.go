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
     
`)
}

func main() {
	if len(os.Args) < 2 {
		PrintUsage()
		return
	}

	// flag parsing
	var help = flag.Bool("help", false, "Codebook is the tool to manage your passcode for all websites.")
	var version = flag.Bool("version", false, "")
	flag.Parse()

	if *help {
		PrintUsage()
		return
	} else if *version {
		fmt.Println(VERSION)
		return
	}

	switch os.Args[1] {
	case "help":
		fmt.Println("TODO(benzh), print all available handlers")
	case "test":
		c := cblib.Init("key")
		c.Add([]byte("baidu.com"), cblib.NewPasscodeHard(15))
		c.Save()
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
			cblib.CopyToClipBoard(string(pc))
		}
	}
}

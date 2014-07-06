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
	"code.google.com/p/gopass"
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
  new <website>             generate a random password for the new website
  add <website>             manually add another entry for the website
  set <website> <password>  manually set the password for a website
  get <website>             return the password for a specific website   
`)
}

func main() {
	if len(os.Args) < 3 {
		PrintUsage()
		return
	}

	if FlagParsing() {
		return
	}

	master_key, err := gopass.GetPass(
		"Enter master key (recommended shorter than 16 bytes):")
	if err != nil {
		panic(err)
	}

	c := cblib.Init(master_key)
	command := os.Args[1]
	website := []byte(os.Args[2])

	switch command {
	case "get":
		if pwd, err := c.Get(website); err == nil {
			fmt.Println(string(pwd))
		} else {
			panic(err)
		}
	case "new":
		var y_or_n string
		if pwd, err := c.Get(website); err == nil {
			fmt.Println(pwd)
			fmt.Println("Password already exists for", string(website))
		} else {
			new_code := cblib.NewPasscodeHard(15)
			fmt.Println(
				"The password for", string(website),
				"is", string(new_code),
				"\nAccept? (y/N):")
			_, _ = fmt.Scanf("%s", &y_or_n)
			if y_or_n != "N" || y_or_n != "n" {
				// when we add, it's encrypted
				c.Add(website, new_code)
				c.Save()
				// only invoke when under OSX environment
				// cblib.CopyToClipBoard(string(new_code))
			}
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

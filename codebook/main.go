package main

import (
	"bytes"
	"code.google.com/p/gopass"
	"codebook/cblib"
	"flag"
	"fmt"
	"os"
)

const (
	VERSION = "0.3"
)

func printUsage() {
	help := "\nUsage: codebook [--version] [--help] <command> [<args>]\n" +
		"\nPassword Management Tools Simplified.\n\nCommands:\n"

	for _, command := range [][]string{
		{"new <website>", "Generate A New Random Password for <website>"},
		{"set <website> <password>", "Set <password> for <website>"},
		{"get <website>", "Get the password for <website>"},
		{"get all", "Get the password for all stored websites"},
	} {
		help += fmt.Sprintf("    %-30.300s%s\n", command[0], command[1])
	}
	fmt.Fprintf(os.Stdout, "%s\n", help)
}

func parseFlags() bool {
	// flag parsing
	var help = flag.Bool("help", false, "Codebook is the tool to manage your passcode for all websites.")
	var version = flag.Bool("version", false, "")

	flag.Parse()

	if *help {
		printUsage()
		return true
	} else if *version {
		fmt.Println(VERSION)
		return true
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	if parseFlags() {
		return
	}

	master_key, err := gopass.GetPass(
		"Enter master key (recommended shorter than 16 bytes): ")
	if err != nil {
		panic(err)
	}

	c := cblib.Init(master_key)
	command := os.Args[1]
	website := []byte(os.Args[2])

	switch command {
	case "get":
		if bytes.Equal(website, []byte("all")) {
			c.PrintPlain()
			return
		}
		if pwd, err := c.Get(website); err == nil {
			fmt.Println(string(pwd))
			cblib.CopyToClipBoard(string(pwd))
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
				cblib.CopyToClipBoard(string(new_code))
			}
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
	wereinit "weremake/init"
)

func main() {
	fmt.Println("Hello, this is weremake")

	flag.Parse()

	args := flag.Args()
	weremakeFile := "weremake.toml"
	if args[0] == "init" {
		if _, err := os.Stat(weremakeFile); os.IsNotExist(err) {
			fmt.Println("No weremake file found, generating one.")
			if len(args) > 1 && len(args[1]) > 0 {
				wereinit.InitToml(args[1])
			} else {
				wereinit.InitToml("MyProject")
			}
		} else {
			fmt.Println("This directory already contains a weremake file. Operation canceled")
			os.Exit(1)
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"os"
	"weremake/werecreate"
)

const (
	weremakeFile = "weremake.toml"
)

func main() {
	flag.Parse()
	args := flag.Args()

	switch args[0] {
	case "init":
		if _, err := os.Stat(weremakeFile); os.IsNotExist(err) {
			fmt.Println("No weremake file found, generating one.")
			if len(args) > 1 && len(args[1]) > 0 {
				werecreate.InitToml(args[1])
			} else {
				werecreate.InitToml("MyProject")
			}
		} else {
			fmt.Println("This directory already contains a weremake file. Operation canceled")
			os.Exit(1)
		}
	case "cmake":
		fmt.Println("awawa")
	}
}

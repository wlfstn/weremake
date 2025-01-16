package main

import (
	"flag"
	"fmt"
	"os"
	wereinit "weremake/init"
)

func main() {
	fmt.Println("Hello, this is weremake")

	initFlag := flag.String("init", "myProject", "Initialize the application")
	flag.Parse()

	weremakeFile := "weremake.toml"
	if *initFlag != "" {
		if _, err := os.Stat(weremakeFile); os.IsNotExist(err) {
			fmt.Println("No weremake file found, generating one.")
			wereinit.InitToml(*initFlag)
		} else {
			fmt.Println("This directory already contains a weremake file. Operation canceled")
			os.Exit(1)
		}
	}
}

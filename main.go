package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, this is weremake")

	initFlag := flag.Bool("init", false, "Initialize the application")
	flag.Parse()

	weremakeFile := "weremake.toml"
	if *initFlag {
		if _, err := os.Stat(weremakeFile); os.IsNotExist(err) {
			fmt.Printf("File '%s' not found. Please ensure it exists before proceeding.\n", weremakeFile)
		}
	}
}

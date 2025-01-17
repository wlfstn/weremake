package werecreate

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type WereMakeSyntax struct {
	PROJECT_NAME  string
	CXX_STANDARD  int8
	SOURCE        []string
	HEADER        []string
	CREATE_STATIC map[string]string
	LINK          map[string][]string
}

func InitToml(prjName string) {

	newWereMake := WereMakeSyntax{
		PROJECT_NAME: prjName,
		CXX_STANDARD: 20,
		SOURCE: []string{
			"main.cpp",
		},
		HEADER:        []string{},
		CREATE_STATIC: map[string]string{},
		LINK:          map[string][]string{},
	}
	file, err := os.Create("weremake.toml")
	if err != nil {
		log.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(newWereMake); err != nil {
		log.Printf("Error encoding TOML: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Succesfully generate weremake file")
}

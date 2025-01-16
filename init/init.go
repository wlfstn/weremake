package wereinit

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type WereMakeSyntax struct {
	ProjectName string
}

func InitToml(prjName string) {

	newWereMake := WereMakeSyntax{
		ProjectName: prjName,
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

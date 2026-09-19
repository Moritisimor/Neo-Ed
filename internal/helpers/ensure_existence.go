package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func IsDir(entry *os.File) bool {
	info, err := entry.Stat()
	if err != nil {
		log.Fatal(err.Error())
	}
	
	return info.IsDir()
}

func EnsureExistence(entryName string) *os.File {
	file, err := os.Open(entryName)
	if err != nil {
		color.PrintBlueln(fmt.Sprintf("'%s' does not exist, so I will create it.", entryName))
	} else {
		return file
	}

	createdFile, creationErr := os.Create(entryName)
	if creationErr != nil {
		log.Fatal(creationErr.Error())
	}

	return createdFile
}
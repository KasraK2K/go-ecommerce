package pkg

import (
	"log"
	"os"
	"path/filepath"
)

func GetRelativePath(path string) string {
	wd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	return filepath.Join(wd, path)
}

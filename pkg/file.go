package pkg

import (
	"os"

	"app/config"
)

func ListOfFileNames(path string) ([]string, error) {
	finalPath := config.AppConfig.UploadPath + "/" + path
	files, err := os.ReadDir(finalPath)
	if err != nil {
		return []string{}, err
	}

	var fileNames []string

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}

func ListOfFileURLs(path string) ([]string, error) {
	finalPath := config.AppConfig.UploadPath + "/" + path
	files, err := os.ReadDir(finalPath)
	if err != nil {
		return []string{}, err
	}

	basePath := config.AppConfig.StaticFileUrl + "/" + path
	var fileNames []string

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, basePath+"/"+file.Name())
		}
	}

	return fileNames, nil
}

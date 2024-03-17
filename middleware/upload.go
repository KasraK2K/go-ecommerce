package middleware

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"

	"app/common"
	"app/config"
	"app/pkg"
)

// TODO : validate file type, validate file size, remove if not valid
func HandleMultipart(c *fiber.Ctx) error {
	contentType := c.Get("Content-Type")

	if strings.Contains(contentType, "multipart/form-data") {
		// If link is valid for upload
		path := string(c.Request().URI().Path())
		isValidPath := false
		for _, item := range config.AppConfig.VALID_UPLOAD_ENDPOINTS {
			if path == item {
				isValidPath = true
			}
		}
		if !isValidPath {
			err := errors.New("upload request is not valid")
			return pkg.JSON(c, err.Error(), http.StatusNotAcceptable)
		}

		status, err := upload(c)
		if err != nil {
			return pkg.JSON(c, err.Error(), status)
		}
		return c.Next()
	}
	return c.Next()
}

func upload(c *fiber.Ctx) (common.Status, error) {
	// Parse the multipart form data
	form, err := c.MultipartForm()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// Get the field name for the files
	mainFolder := form.Value["m"][0]
	subFolder := form.Value["s"][0]

	// Create a new directory based on the field name
	dirPath := filepath.Join("uploads", mainFolder, subFolder)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return http.StatusInternalServerError, err
	}

	// Loop through all the files in the form data
	for fileFieldName, fileHeader := range form.File {
		// Loop through all the files for this field name
		for _, file := range fileHeader {
			fileExt := filepath.Ext(file.Filename)
			if fileExt == ".jpeg" {
				fileExt = ".jpg"
			}

			// Check file type
			isExtValid := false
			for _, ext := range config.AppConfig.FILE_EXTENSIONS {
				if ext == fileExt {
					isExtValid = true
				}
			}
			if !isExtValid {
				return http.StatusNotAcceptable, errors.New("file extension is not valid")
			}

			// Check File Size
			if file.Size > config.AppConfig.FILE_SIZE {
				return http.StatusNotAcceptable, errors.New("file is too large")
			}

			// Open the file
			fileContent, err := file.Open()
			if err != nil {
				return http.StatusInternalServerError, err
			}
			defer fileContent.Close()

			// Create a new file on the server
			newFileName := filepath.Join(dirPath, fileFieldName+fileExt)
			newFile, err := os.Create(newFileName)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			defer newFile.Close()

			// Write the file to the server
			if _, err := io.Copy(newFile, fileContent); err != nil {
				return http.StatusInternalServerError, err
			}
		}
	}

	return http.StatusOK, nil
}

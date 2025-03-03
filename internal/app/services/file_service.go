package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func FileSave(c *gin.Context, fileField string, uploadDir string) (string, error) {
	file, err := c.FormFile(fileField)
	if err != nil {
		return "", fmt.Errorf("file not found: %w", err)
	}

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	filePath := filepath.Join(uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	return file.Filename, nil
}

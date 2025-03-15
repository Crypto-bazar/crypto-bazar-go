package utils

import "fmt"

func GenerateTokenURI(baseUrl string, image string) string {
	return fmt.Sprintf("%s/metadata/%s", baseUrl, image)
}
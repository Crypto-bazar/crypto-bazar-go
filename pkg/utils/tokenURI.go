package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateTokenURI(baseUrl string) string {
	id := uuid.New().String() // например, "3e4666bf-d5e5-4aa7-b8ce-cefe41c7568a"
	return fmt.Sprintf("%s/metadata/%s", baseUrl, id)
}

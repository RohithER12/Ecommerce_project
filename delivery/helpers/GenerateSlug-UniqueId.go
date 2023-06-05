package helpers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

func GenerateUniqueID(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("invalid length")
	}

	numBytes := length / 2
	if length%2 != 0 {
		numBytes++
	}

	bytes := make([]byte, numBytes)
	_, err := rand.Read(bytes)
	if err != nil {

		return "", errors.New("failed to create unique id")
	}

	uniqueID := base64.URLEncoding.EncodeToString(bytes)
	fmt.Println("\n uniqueID", uniqueID)

	return uniqueID, nil
}

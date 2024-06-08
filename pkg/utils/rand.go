package utils

import (
	"crypto/rand"
	"encoding/base64"
)


func RandomString(length int) (string, error){
	bytes := make([]byte, length) 
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
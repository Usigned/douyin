package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.NewString()
}

func Md5(key string) string {
	data := []byte(key)
	h := md5.Sum(data)
	return fmt.Sprintf("%x", h)
}

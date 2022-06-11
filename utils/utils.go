package utils

import (
	"crypto/md5"
	"encoding/hex"
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

func EncryptPassword(key string) string {
	data := []byte(key)
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(data))
}

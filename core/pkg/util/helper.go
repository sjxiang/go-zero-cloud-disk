package util

import (
	"crypto/md5"
	"fmt"
)

func MD5(plainText string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(plainText)))
}
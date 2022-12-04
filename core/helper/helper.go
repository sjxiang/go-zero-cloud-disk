package helper

import (
	"crypto/md5"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func MD5(plainText string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(plainText)))
}

func GenerateToken(id uint64, identity, name string) (string, error) {
	uc := CustomUserClaim {
		Id: id,
		Identity: identity,
		Name: name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
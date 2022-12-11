package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomUserClaim struct {
	Id       uint64
	Identity string
	Name     string

	jwt.StandardClaims
}


var JWTSecretKey = "cnd1-24enfilvbib"


func GenerateToken(id uint64, identity, name string, second int64) (string, error) {
	uc := CustomUserClaim {
		Id: id,
		Identity: identity,
		Name: name,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix(),  // 有效时间
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func AnalyzeToken(token string) (*CustomUserClaim, error) {
	uc := new(CustomUserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims.Valid {
		return uc, errors.New("token is invalid")
	}

	return uc, nil
}
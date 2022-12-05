package jwt

import "github.com/dgrijalva/jwt-go"

type CustomUserClaim struct {
	Id       uint64
	Identity string
	Name     string

	jwt.StandardClaims
}


var JWTSecretKey = "cnd1-24enfilvbib"


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

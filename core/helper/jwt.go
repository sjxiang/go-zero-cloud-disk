package helper

import "github.com/dgrijalva/jwt-go"

type CustomUserClaim struct {
	Id       uint64
	Identity string
	Name     string

	jwt.StandardClaims
}


var JWTSecretKey = "cnd1-24enfilvbib"
package model

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID          uint
	UserName    string
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}

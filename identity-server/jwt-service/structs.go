package jwtservice

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	ID       int64
	Username string
	Password string
	jwt.StandardClaims
}

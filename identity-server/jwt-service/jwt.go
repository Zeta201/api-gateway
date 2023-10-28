package jwtservice

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/Zeta201/identity-server/model"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
}

func NewAccessToken(user model.User) (string, error) {
	key := os.Getenv("KEY")
	var claims = UserClaims{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return accessToken.SignedString([]byte(key))
}

// func NewRefreshToken(claims jwt.StandardClaims, key string) (string, error) {
// 	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return refreshToken.SignedString([]byte(key))
// }

func ParseAccessToken(accessToken string) (*UserClaims, error) {
	key := os.Getenv("KEY")
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedAccessToken.Valid {
		return nil, errors.New("invalid token")
	}
	claims, ok := parsedAccessToken.Claims.(*UserClaims)

	if !ok {
		return nil, errors.New("ID token valid but could not parse claims")
	}

	return claims, nil
}

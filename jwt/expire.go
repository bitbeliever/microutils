package jwtutil

import (
	"github.com/golang-jwt/jwt"
	"time"
)

func IsExpire(tokenStr string) bool {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return true
	}
	exp, ok := token.Header["exp"].(int64)
	if !ok {
		return true
	}

	expireAt := time.Unix(exp, 0)
	if time.Now().After(expireAt) {
		return true
	}

	return false
}

package utils

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var hmacSampleSecret []byte

func init() {
	hmacSampleSecret = []byte("bitway-todo_block")
}

func CreateToken() (string, error) {
	timeByte, _ := json.Marshal(time.Now().Unix())
	jti := fmt.Sprintf("%x", md5.Sum(timeByte)) //将[]byte转成16进制

	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"token": jti,
		"exp":   time.Now().Add(time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	})

	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func GetVerifyToken(tokenString string) (string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if err != nil {
		return "", false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", false
	}

	return claims["token"].(string), true
}

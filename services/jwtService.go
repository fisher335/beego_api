package services

import (
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

func CreateToken(Phone string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	tokenexpConfig, _ := config.String("Tokenexp")
	tokenexp, _ := strconv.Atoi(tokenexpConfig)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenexp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["phone"] = Phone
	token.Claims = claims
	TokenSecrets, _ := config.String("TokenSecrets")
	tokenString, _ := token.SignedString([]byte(TokenSecrets))
	return tokenString
}

func CheckToken(tokenString string) (bool, string) {
	Phone := ""
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		TokenSecrets, _ := config.String("TokenSecrets")
		return []byte(TokenSecrets), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	Phone = claims["phone"].(string)
	exp := claims["exp"]
	if exp.(int64) < time.Now().Unix() {
		return false, Phone
	}
	return true, Phone
}

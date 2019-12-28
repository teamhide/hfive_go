package utils

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

type TokenHandler struct{}

func (t TokenHandler) EncodeToken(payload map[string]string) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		payload,
        "exp":  time.Now().Add(time.Hour * time.Duration(1)).Unix(),
    })
}
func (t TokenHandler) DecodeToken(token string) map[string]string {
	temp := make(map[string]string)
	return temp
}

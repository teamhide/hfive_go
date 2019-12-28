package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Util struct{}

func (u Util) EncodeToken(payload map[string]string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		payload,
		"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})
}
func (u Util) DecodeToken(token string) map[string]string {
	temp := make(map[string]string)
	return temp
}

func (u Util) 
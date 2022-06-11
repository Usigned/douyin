package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type MyClaims struct {
	Username           string `json:"id,omitempty"`
	Password           string `json:"password"`
	jwt.StandardClaims        // 官方包含的字段
}

// 定义secret
var mySecret = []byte("中国地质大学")

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

// TokenExpireDuration 定义JWT的过期时间 ..
const TokenExpireDuration = time.Hour * 24

// GenToken 生成Token
func GenToken(username, password string) (Token string, err error) {
	c := MyClaims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration * 30).Unix(), // 过期时间
			Issuer:    "pcvocaloid",                                    // 签发人
		},
	}
	Token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Token", Token)
	return
}

func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = &MyClaims{}
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid {
		err = Error{Msg: "invalid token"}
	}
	return
}

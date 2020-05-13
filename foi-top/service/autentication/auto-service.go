package aut_service

import (
	"github.com/Andreson/go-portifolio/foi-top/domain"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("o bico do peito do chico preto e preto e quem disser q o bico do peito do chico preto nao e preto tem o bico do peiro mais preto q o bico do peito do chico preto")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(login domain.LoginDTO) (string, error){

	claims := &Claims{
		Username: login.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationDateToken() } }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if tokenString, err := token.SignedString(jwtKey); err==nil{
		return tokenString, nil
	} else {return "", err}
}

func ValidatedToken(login domain.LoginDTO) (bool, error){

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(login.TokenJWT, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return false, err
	} else {return token.Valid, nil}
}

 func RewToken(login domain.LoginDTO)(string, error){

	 claims := &Claims{}
	 _, err := jwt.ParseWithClaims(login.TokenJWT, claims, func(token *jwt.Token) (interface{}, error) {
		 return jwtKey, nil
	 })

	 if err != nil {  return "", err }
	 token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	 tokenString, err := token.SignedString(jwtKey)
	 if err != nil {
		 return "", nil
	 } else {return tokenString, nil }
 }

 func expirationDateToken()int64 {
	return time.Now().Add(50000 * time.Minute).Unix()

 }
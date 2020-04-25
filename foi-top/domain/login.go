package domain


import 	"github.com/dgrijalva/jwt-go"

type Login struct {
	Username string
	Password string
	Enable bool
	TokenJWT string
	StandardClaims jwt.StandardClaims
	User Usuario
}

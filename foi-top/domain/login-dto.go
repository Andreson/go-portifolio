package domain


import 	"github.com/dgrijalva/jwt-go"

type LoginDTO struct {
	Username       string
	Password       string
	Enable         bool
	TokenJWT       string
	StandardClaims jwt.StandardClaims
	UserId int64

}

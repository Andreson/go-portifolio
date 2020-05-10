package domain


import 	"github.com/dgrijalva/jwt-go"

type LoginDto struct {
	Username       string
	Password       string
	Enable         bool
	TokenJWT       string
	StandardClaims jwt.StandardClaims
	User           UserDTO
}

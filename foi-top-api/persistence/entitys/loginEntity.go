package event_entity

import "github.com/dgrijalva/jwt-go"

type LoginEntity struct {

	Username       string
	Password       string
	Enable         bool
	TokenJWT       string
	StandardClaims jwt.StandardClaims
	User  UsuarioEntity `gorm:"foreignkey:UserID"`
	UserID int64
}


func (LoginEntity) TableName()string {
	return "login"
}
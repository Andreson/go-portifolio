package login_service

import (
	"crypto/sha1"
	"encoding/base64"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	login_dao "github.com/Andreson/go-portifolio/foi-top/persistence/dao/login"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
)

func Save(login domain.LoginDTO){
	login_dao.Salvar(ToEntity(login))
}

func FindByUser(userID int64) domain.LoginDTO {
	login :=login_dao.FindByUserID(userID)

	return ToDTO(login)
}

func ToDTO(login event_entity.LoginEntity)domain.LoginDTO{


	return domain.LoginDTO{
		Enable:login.Enable,
		Password:"##########",
		Username:login.Username,
		UserId:login.UserID }
}

func ToEntity(login domain.LoginDTO)event_entity.LoginEntity{

	sh :=sha1.New()
	sh.Write([]byte(login.Password))
	pass :=base64.URLEncoding.EncodeToString(sh.Sum(nil) )


   return event_entity.LoginEntity{
		Enable:true,
		Password:pass,
		Username:login.Username,
		UserID:login.UserId }
}
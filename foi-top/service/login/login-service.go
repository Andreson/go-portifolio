package login_service

import (
	"crypto/sha1"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	login_dao "github.com/Andreson/go-portifolio/foi-top/persistence/dao/login"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
)

func Save(login domain.LoginDTO){
	login_dao.Salvar(ToEntity(login))
}


func ToEntity(login domain.LoginDTO)event_entity.LoginEntity{

	sh :=sha1.New()
	sh.Write([]byte(login.Password))
	pass :=string(sh.Sum(nil))

   return event_entity.LoginEntity{
		Enable:true,
		Password:pass,
		Username:login.Username,
		UserID:login.UserId }
}
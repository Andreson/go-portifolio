package user_service

import (
	"github.com/Andreson/go-portifolio/foi-top/domain"
	user_dao "github.com/Andreson/go-portifolio/foi-top/persistence/dao/usuario"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	login_service "github.com/Andreson/go-portifolio/foi-top/service/login"
	"github.com/dgrijalva/jwt-go"
	"log"
)
func Save(user domain.UserDTO) {

   userSave := user_dao.Salvar(ToEntity(user))

	loginDTO := domain.LoginDTO{
		Username:       user.Login.Username,
		Password:       user.Login.Password,
		Enable:         true,
		TokenJWT:       "",
		StandardClaims: jwt.StandardClaims{},
		UserId:         userSave.ID,
	}

	login_service.Save(loginDTO)
   log.Println("new user ", userSave)

}


func SendInviteEvent(invite domain.InviteEventDTO){

}


func FindById(user domain.UserDTO)domain.UserDTO{
	userDto := ToDTO(user_dao.FindById(user.Id) )
	userDto.Login=login_service.FindByUser(userDto.Id)
	return userDto;
}

func ListByEvent(eventoId int64) {

}


func ToDTO(user event_entity.UsuarioEntity) domain.UserDTO {
	return domain.UserDTO {
		Nome:user.Nome,
		Endereco:user.Endereco,
		Email:user.Email,
		Fone:user.Fone,
		Id:user.ID,

	}
}

func ToEntity(user domain.UserDTO) event_entity.UsuarioEntity {
	return event_entity.UsuarioEntity {
		Nome:user.Nome,
		Endereco:user.Endereco,
		Email:user.Email,
		Fone:user.Fone,

	}
}
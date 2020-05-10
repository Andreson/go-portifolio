package user_service

import (
	"github.com/Andreson/go-portifolio/foi-top/domain"
	user_dao "github.com/Andreson/go-portifolio/foi-top/persistence/dao/usuario"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"log"
)

func Save(user domain.UserDTO) {
   userSave := user_dao.Salvar(ToEntity(user))

   log.Println("new user ", userSave)

}


func FindById(user domain.UserDTO)domain.UserDTO{
	userEntity := user_dao.FindById(user.Id)
	return ToDTO(userEntity)
}

func ListByEvent(eventoId int64) {

}


func ToDTO(user event_entity.UsuarioEntity) domain.UserDTO {
	return domain.UserDTO {
		Nome:user.Nome,
		Endereco:user.Endereco,
		Email:user.Email,
		Fone:user.Fone,

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
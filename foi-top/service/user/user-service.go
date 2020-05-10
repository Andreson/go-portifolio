package user_service

import (
	"github.com/Andreson/go-portifolio/foi-top/domain"
	user_dao "github.com/Andreson/go-portifolio/foi-top/persistence/dao/usuario"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
)

func Save(user domain.UserDTO) {
  user_dao.Salvar(ToEntity(user))

}


func FindById(user domain.UserDTO){
	user_dao.FindById(user.Id)
}

func ListByEvent(eventoId int64) {

}

func ToEntity(user domain.UserDTO) event_entity.UsuarioEntity {
	return event_entity.UsuarioEntity {
		Nome:user.Nome,
		Endereco:user.Endereco,
		Email:user.Email,
		Fone:user.Fone,

	}
}
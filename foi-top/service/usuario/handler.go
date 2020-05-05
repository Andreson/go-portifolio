package user_service

import (
	"github.com/Andreson/go-portifolio/foi-top/domain"
	user_dao "github.com/Andreson/go-portifolio/foi-top/persistence/dao/usuario"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
)

func Cadastrar(user domain.UsuarioDTO) {
  user_dao.Salvar(ToEntity(user))
}


func ToEntity(user domain.UsuarioDTO) event_entity.UsuarioEntity {
	return event_entity.UsuarioEntity {
		Nome:user.Nome,
		Endereco:user.Endereco,
		Email:user.Email,
		Fone:user.Fone,
	}
}
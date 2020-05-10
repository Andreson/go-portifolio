package user_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/persistence"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
)

func Salvar(user event_entity.UsuarioEntity)event_entity.UsuarioEntity {
 persistence.Execute(func(db *gorm.DB) *gorm.DB {
		return db.Save(&user)
	})
  return user
}


func FindById(userId int64){

}

func ListarUsuarioPorEvento(idEvento int64) {

}
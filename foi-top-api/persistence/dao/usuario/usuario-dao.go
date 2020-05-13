package user_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/config"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
)

func Salvar(user event_entity.UsuarioEntity)event_entity.UsuarioEntity {
 config.ExecuteQuery(func(db *gorm.DB) *gorm.DB {
		return db.Save(&user)
	})
  return user
}

func FindById(userId int64)event_entity.UsuarioEntity{
	var user event_entity.UsuarioEntity
	config.ExecuteQuery(func(db *gorm.DB) *gorm.DB {
		return db.First(&user,userId)
	})
	return user
}

func ListarUsuarioPorEvento(idEvento int64) {

}
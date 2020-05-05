package user_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/persistence"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
)

func Salvar(user event_entity.UsuarioEntity) {
	persistence.Execute(func(db *gorm.DB) error {
		return	db.Save(user).Error
	})
}
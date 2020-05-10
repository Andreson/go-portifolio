package login_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/persistence"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
)

func Salvar(user event_entity.LoginEntity)event_entity.LoginEntity {
	resp :=persistence.Execute(func(db *gorm.DB) *gorm.DB {
		return db.Save(user)
	})
	return resp.Value.(event_entity.LoginEntity)
}

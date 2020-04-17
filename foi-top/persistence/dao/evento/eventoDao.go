package evento_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/persistence"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
)

func Save(entity event_entity.EventoEntity ){
	persistence.Execute(func(db *gorm.DB) {
		db.Create(&entity)
	})
}

func FindById(entity event_entity.EventoEntity )[]event_entity.EventoEntity{
	var response []event_entity.EventoEntity
	persistence.Execute(func(db *gorm.DB) {
		db.First(&response, entity.ID)
	})
	return response
}

package evento_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/persistence"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
	"log"
)

func Save(entity event_entity.EventoEntity ){
	persistence.Execute(func(db *gorm.DB)*gorm.DB  {
		return db.Create(&entity)
	})
}

func FindById(entity event_entity.EventoEntity )event_entity.EventoEntity{
	var response event_entity.EventoEntity
	persistence.Execute(func(db *gorm.DB)*gorm.DB  {
		return db.First(&response, entity.ID)
	})

	log.Println("Evento procurado ", response)
	return response
}

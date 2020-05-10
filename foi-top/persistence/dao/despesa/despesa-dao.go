package despesa_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/persistence"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
)

func Save(entity event_entity.ItemDespesaEventoEntity ){
	persistence.Execute(func(db *gorm.DB)*gorm.DB {
		return db.Create(&entity)
	})
}

func ListByEvent(eventId int,response *[]event_entity.ItemDespesaEventoEntity){
	persistence.Execute(func(db *gorm.DB)*gorm.DB  {
		resp:=db.Where("evento_refer =?",eventId).Find(response)
		return resp
	})
}

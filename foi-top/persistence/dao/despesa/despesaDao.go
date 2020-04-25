package despesa_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/persistence"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
)

func Save(entity event_entity.ItemDespesaEventoEntity ){
	persistence.Execute(func(db *gorm.DB)error {
		return db.Create(&entity).Error
	})
}

func ListByEvent(eventId int,response *[]event_entity.ItemDespesaEventoEntity){
	persistence.Execute(func(db *gorm.DB)error {
		resp:=db.Where("evento_refer =?",eventId).Find(response)
		return resp.Error
	})
}

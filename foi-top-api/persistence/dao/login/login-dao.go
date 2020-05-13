package login_dao

import (
	"github.com/Andreson/go-portifolio/foi-top/config"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
)

func Salvar(userLogin event_entity.LoginEntity)event_entity.LoginEntity {
	 config.ExecuteQuery(func(db *gorm.DB) *gorm.DB {
		return db.Save(&userLogin)
	})
	return userLogin
}


func FindByUserID(userId int64)event_entity.LoginEntity{
	var temp event_entity.LoginEntity

	config.ExecuteQuery(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", userId).Find(&temp)
	})
	return temp
}
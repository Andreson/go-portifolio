package persistence

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


var db gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:123@/eventos?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		log.Panic("Erro ao inicializar conexao com banco  de dados ",err)
		defer db.Close()
	}
}

func GetConn() gorm.DB {
	return db
}

func Save(entity interface{} ){

	db.Create(&entity)
	defer db.Close()
}

func FindAll(entity interface{}) *gorm.DB {
 return  db.Find(&entity)
}

func FindById(id int,entity interface{} )interface{} {
	return db.First(&entity,id)
}
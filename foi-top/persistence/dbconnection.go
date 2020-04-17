package persistence

import (
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


var db gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:123@/eventos?charset=utf8&parseTime=True&loc=Local")
	db.Table("evento").CreateTable(&event_entity.EventoEntity{})
	if err!=nil {
		log.Panic("Erro ao inicializar conexao com banco  de dados ",err)
		defer db.Close()
	}
}

func Execute( exec func(*gorm.DB) ) {
			db:=GetConn();
			exec(db)
			defer  db.Close();
}

func GetConn() *gorm.DB {

	db, err := gorm.Open("mysql", "root:123@/eventos?charset=utf8&parseTime=True&loc=Local")
	db.Table("evento").CreateTable(&event_entity.EventoEntity{})
	if err!=nil {
		log.Panic("Erro ao inicializar conexao com banco  de dados ",err)
		defer db.Close()
	}

	return db
}

func Save(entity event_entity.EventoEntity ){
	db, _ := gorm.Open("mysql", "root:123@/eventos?charset=utf8&parseTime=True&loc=Local")
	db.Create(&entity)
	defer db.Close()
}

func FindAll(entity interface{}) *gorm.DB {
 return  db.Find(&entity)
}

func FindById(id int,entity interface{} )interface{} {
	return db.First(&entity,id)
}
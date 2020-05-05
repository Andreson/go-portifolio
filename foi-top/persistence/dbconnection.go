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
	db.AutoMigrate(&event_entity.EventoEntity{},&event_entity.ItemDespesaEventoEntity{})

	if err!=nil {
		log.Panic("Erro ao inicializar conexao com banco  de dados ",err)
		defer db.Close()
	}
}

func Execute( exec func(*gorm.DB) error )error {
			db:=GetConn()
			 if isError := exec(db); isError!=nil{
			 	log.Println("Ocorreu um erro nao esperado ao executar  query: ",isError)
			 	return isError
			 }else {
				log.Println("Objeto de erro retornado  ", isError)
			}
			defer  db.Close()
			return nil
}

func GetConn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123@/eventos?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		log.Panic("Erro ao inicializar conexao com banco  de dados ",err)
		defer db.Close()
	}

	return db
}



func FindAll(entity interface{}) *gorm.DB {
 return  db.Find(&entity)
}

func FindById(id int,entity interface{} )interface{} {
	return db.First(&entity,id)
}
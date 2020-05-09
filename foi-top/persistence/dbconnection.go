package persistence

import (
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)


var db *gorm.DB

func init() {
	GetConn()
	db.AutoMigrate(&event_entity.EventoEntity{},&event_entity.ItemDespesaEventoEntity{})


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
	var  profile  string
	var err error
	 if profile =os.Getenv("PROFILE"); profile=="" {
	 	profile="dev"
	 }
	 if profile=="dev" {
		 db, err = gorm.Open("mysql", "root:123@/eventos?charset=utf8&parseTime=True&loc=Local")
	 } else {
		 db, err = gorm.Open("mysql", "root:123123@(104.198.32.149)/eventos?charset=utf8&parseTime=True&loc=Local")
	 }
	if err!=nil {
		log.Printf("\nPROFILE %s",os.Getenv("PROFILE"))
		log.Panic("Erro ao inicializar conexao com banco  de dados , profile  ",err,profile)
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
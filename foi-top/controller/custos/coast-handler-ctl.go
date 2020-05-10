package coast_controller

import (
	"encoding/json"
	. "github.com/Andreson/go-portifolio/foi-top/controller"
	aut_controller "github.com/Andreson/go-portifolio/foi-top/controller/autentication"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	itemevento_service "github.com/Andreson/go-portifolio/foi-top/service/itemEvento"
	"log"
	"net/http"
	"strconv"
	"strings"
)


func Init(){

	initializeFuncHandler()
}

func AddCoast(w http.ResponseWriter, r *http.Request) {
	data,err :=toDto(r)
	DefaultCall( func(){
		itemevento_service.Cadastrar(data)
	} ,w,err)
}

func ListCoastByEvent(w http.ResponseWriter, r *http.Request) {

	tempIdParam := strings.TrimPrefix(r.URL.Path, "/event/coast/")
	idEvent, err :=strconv.Atoi(tempIdParam)
	DefaultCallBody( func()interface{}{
		return itemevento_service.ListByEvent(idEvent)
	} ,w,err)
}

func  toDto(r *http.Request) (domain.ItemCoastEventDTO,error) {
	decoder := json.NewDecoder(r.Body)
	var data domain.ItemCoastEventDTO
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	log.Print("Request ",data)
	return data, nil
}


func initializeFuncHandler() {

	addCoast := http.HandlerFunc(AddCoast)
	http.Handle("/event/coast",aut_controller.FilterHandler(addCoast))

	listCoastByEvent := http.HandlerFunc(ListCoastByEvent)
	http.HandleFunc("/event/coast/", listCoastByEvent)
}
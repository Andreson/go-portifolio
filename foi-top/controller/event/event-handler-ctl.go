package event_controller

import (
	"encoding/json"
	"github.com/Andreson/go-portifolio/foi-top/controller"
	aut_controller "github.com/Andreson/go-portifolio/foi-top/controller/autentication"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	event_service "github.com/Andreson/go-portifolio/foi-top/service/event"
	"log"
	"net/http"
	"strconv"
	"strings"
)


func Init(){
	initializeFuncHandler()
}

func Create(w  http.ResponseWriter, r *http.Request){
	dto, err :=toDto(r)
	controller.DefaultCall( func(){
		event_service.Save(dto)
	} ,w,err)
}

func FindById(w  http.ResponseWriter, r *http.Request) {

	tempIdParam := strings.TrimPrefix(r.URL.Path, "/event/")
	id, err :=strconv.Atoi(tempIdParam)

	controller.DefaultCallBody( func()interface{}{
		return event_service.FindById(domain.EventDto{Id: id})
	} ,w,err)

}

func  toDto(r *http.Request) (domain.EventDto,error) {
	decoder := json.NewDecoder(r.Body)
	var data domain.EventDto
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	 log.Print("Request ",data)
	   return data, nil
}

func initializeFuncHandler(){
	createEvent := http.HandlerFunc(Create)
	http.Handle("/event",aut_controller.FilterHandler(createEvent))

	findById:= http.HandlerFunc(FindById)
	http.Handle("/event/",aut_controller.FilterHandler(findById))

}




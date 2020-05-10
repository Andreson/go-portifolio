package user_controller

import (
	"encoding/json"
	"github.com/Andreson/go-portifolio/foi-top/controller"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	event_service "github.com/Andreson/go-portifolio/foi-top/service/evento"
	"log"
	"net/http"
)

func Save(w  http.ResponseWriter, r *http.Request){

	dto, err :=toDto(r)
	controller.DefaultCall( func(){
		event_service.Save(dto)
	} ,w,err)
}

func  toDto(r *http.Request) (domain.ItemDespesaEventoDTO,error) {
	decoder := json.NewDecoder(r.Body)
	var data domain.ItemDespesaEventoDTO
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	log.Print("Request ",data)
	return data, nil
}

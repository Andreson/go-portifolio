package event_controller

import (
	"encoding/json"
	"github.com/Andreson/go-portifolio/foi-top/controller"
	event_model "github.com/Andreson/go-portifolio/foi-top/model/evento"
	event_service "github.com/Andreson/go-portifolio/foi-top/service/evento"
	"log"
	"net/http"
)

func Create(w  http.ResponseWriter, r *http.Request){

	dto, erro :=toDto(r)
	w.Header().Set("Content-Type", "application/json")
	if erro==nil {
		event_service.Cadastrar(dto)
		w.Write(controller.GetSuccess())
	 }else {
	 	  w.WriteHeader(400)
	 	w.Write( controller.GetError(erro.Error()))
	}
}

func FindById(w  http.ResponseWriter, r *http.Request) {

}


func  toDto(r *http.Request) (event_model.EventoDto,error) {
	decoder := json.NewDecoder(r.Body)
	var data event_model.EventoDto
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	 log.Print("Request ",data)
	   return data, nil
}




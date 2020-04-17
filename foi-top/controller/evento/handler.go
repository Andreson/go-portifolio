package event_controller

import (
	"encoding/json"
	"github.com/Andreson/go-portifolio/foi-top/controller"
	event_model "github.com/Andreson/go-portifolio/foi-top/model/evento"
	user_model "github.com/Andreson/go-portifolio/foi-top/model/usuario"
	event_service "github.com/Andreson/go-portifolio/foi-top/service/evento"
	"log"
	"net/http"
	"strconv"
)

func Create(w  http.ResponseWriter, r *http.Request){

	dto, erro :=toDto(r)
	w.Header().Set("Content-Type", "application/json")
	if erro==nil {
		event_service.Cadastrar(dto)
	 }else {
	 	  w.WriteHeader(400)
	 	 resp,err:=json.Marshal(controller.ReponseModel{StatusCode:400,Message:"Ocorreu um erro ao converter entidade," +erro.Error()})
		if err==nil{
			w.Write(resp)
		}
	 	w.Write([]byte("Erro nao esperado ao converter meg de erro "))
	}




}



func  toDto(r *http.Request) (event_model.EventoDto,error) {
	decoder := json.NewDecoder(r.Body)
	var data event_model.EventoDto
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	owner := data.Data


	log.Println("Data recebida ",owner)

	if err!=nil {
		log.Println("Erro ao converter data do evento ",err)
		return event_model.EventoDto{},err
	}
	userID :=r.Form.Get("userid")
	idUser, userErr:= strconv.Atoi(userID)
	if userErr!=nil {
		log.Println("Erro ao converter parametro Userid",userErr)
		return event_model.EventoDto{},userErr
	}
	return event_model.EventoDto{Titulo: r.Form.Get("titulo"),
		Endereco:r.Form.Get("endereco"),
		Bairro: r.Form.Get("bairro"),
		Usuario:user_model.Usuario{Id: idUser } ,
		Data: "time.Now()" }, nil
}




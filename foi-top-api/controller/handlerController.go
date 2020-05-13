package controller

import (
	"encoding/json"
	"net/http"

	"log"
)

type ReponseModel  struct {
	Message string `json:"message"`
	Body interface{} `json:"response"`

}

func InitEndPoints() {

}


//faz tratamento de erro e redirecimamento padrao das buscas
func DefaultCallBody(handler func() interface{},w http.ResponseWriter, err error){
	w.Header().Set("Content-Type", "application/json")
	if err==nil {
		response :=handler()
		w.Write(GetSuccessBody(response))
	}else {
		w.WriteHeader(400)
		w.Write( GetError(err.Error()))
	}

}

//faz tratamento de erro e redirecimamento padrao das buscas
func DefaultCall(handler func(),w http.ResponseWriter, err error){
	w.Header().Set("Content-Type", "application/json")
	if err==nil {
		handler()
		w.Write(GetSuccess())
	}else {
		w.WriteHeader(400)
		w.Write( GetError(err.Error()))
	}

}

func GetSuccessBody(body interface{}) []byte{
	if resp, err :=json.Marshal(ReponseModel{Message:"Operaçao realizado com sucesso",Body:body}) ; err==nil {
		return resp
	}else {
		log.Println("Ocorreu um erro nao esperado ao fazer Marshal GetSuccess ",err)
		return defaulError()
	}
}

func GetSuccessMsg(msg string) []byte{
	if resp, err :=json.Marshal(ReponseModel{Message:msg}) ; err==nil {
		return resp
	}else {
		log.Println("Ocorreu um erro nao esperado  ",err)
		return defaulError()
	}
}

func GetSuccess() []byte{
	if resp, err :=json.Marshal(ReponseModel{Message:"Operaçao realizado com sucesso"}) ; err==nil {
		return resp
	}else {
		log.Println("Ocorreu um erro nao esperado ao fazer Marshal GetSuccess ",err)
		return defaulError()
	}
}

func GetError(msg string) ([]byte){
	if resp, err := json.Marshal(ReponseModel{Message:"Ocorreu um erro ao converter entidade," +msg}) ;err==nil {
		return resp
	} else {
		log.Println("Ocorreu um erro nao esperado ao fazer Marshal GetError ",err)
		return defaulError()
	}
}

func defaulError()([]byte){
	 return []byte("Ocorreu um erro nao esperando")
}
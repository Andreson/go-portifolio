package controller

import (
	"encoding/json"

	"log"
)

type ReponseModel  struct {
	Message string
	Body interface{}
}

func GetSuccessBody(body interface{}) []byte{
	if resp, err :=json.Marshal(ReponseModel{Message:"Operaçao realizado com sucesso",Body:body}) ; err==nil {
		return resp
	}else {
		log.Println("Ocorreu um erro nao esperado ao fazer Marshal GetSuccess ",err)
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
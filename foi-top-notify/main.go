package main

import (
	"encoding/json"
	"fmt"
	notify_service "github.com/Andreson/go-portifolio/foi-top-notify/service"
)

/**
Programa destinado a estudos em golang
componente de sistema q le mensagens do pubsub da google e  executa as notificaçoes

 */
func main(){


	resp, err:=notify_service.ReadMesage("sub-send-invite-event")

	if err!=nil{
		fmt.Println("err s",err)
	}

	var teste Test

	json.Unmarshal(resp , &teste)

	fmt.Println("iae",teste)

}

type Test struct {
	Nome string
	Idade int
}
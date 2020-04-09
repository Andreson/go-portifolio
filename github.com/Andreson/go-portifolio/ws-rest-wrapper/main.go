package main

import (
	"fmt"
	"github.com/Andreson/go-portifolio/ws-rest-wrapper/client-ws"
)

func main(){
	c := client_ws.RequestTemplate{Url: "https://enpopwf3y505.x.pipedream.net/", Body :Book{Id:434,Nome:"O diario de Any Frank"}}
	var response = Resposta{}
	c.Post().Body(&response)
	fmt.Println("Resposta do serviço ",response.Success)
}


type Resposta struct {
	Success bool `json:"success"`
}

type Book struct {
	Id   int    `json:"userId"`
	Nome string `json:"title"`
}
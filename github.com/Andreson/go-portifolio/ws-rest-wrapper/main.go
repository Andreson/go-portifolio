package main

import (
	"fmt"
	"github.com/Andreson/go-portifolio/ws-rest-wrapper/client-ws"
)

func main(){
	c := client_ws.RequestTemplate{Url: "https://jsonplaceholder.typicode.com/"}
	c.AddHeader(client_ws.RequestHeader{Name: "user-id",Value:"123456"})
	var response = Book{}

	 c.Get("todos/1").Body(response)
	fmt.Println("Resposta do serviço ",response)
}

type Book struct {
	Id   int    `json:"userId"`
	Nome string `json:"title"`
}
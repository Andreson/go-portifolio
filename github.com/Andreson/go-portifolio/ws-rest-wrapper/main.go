package main

import (
	"fmt"
	client_ws "github.com/Andreson/go-portifolio/ws-rest-wrapper/client-ws"
)

func main(){
	respoModel :=Book{}
	c := client_ws.Request{Url:"https://jsonplaceholder.typicode.com/"}
	c.Get("todos/1",&respoModel)
	fmt.Println("Reposta do serviço ",respoModel)
}

type Book struct {
	Id   int    `json:"userId"`
	Nome string `json:"title"`
}
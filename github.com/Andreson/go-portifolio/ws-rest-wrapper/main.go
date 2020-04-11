package main

import (
	"fmt"
	client_ws "github.com/Andreson/go-portifolio/ws-rest-wrapper/client-ws"
)

func main(){


}

//Exemplo de utilizaçao dos pacotes para fazer um request simples com o metodo http Get
func GetExample(){

	c := client_ws.RequestTemplate{Url: "https://jsonplaceholder.typicode.com/"}
	c.AddHeader(client_ws.RequestHeader{Name: "user-id",Value:"123456"})
	var response = Book{}

	c.Get("todos/1").Body(response)

}

//Exemplo de utilizaçao dos pacotes para fazer um request simples com o metodo http Post
func PostExample(){

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
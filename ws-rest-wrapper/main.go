package main

import (
	"fmt"
	client_ws "github.com/Andreson/go-portifolio/ws-rest-wrapper/client-ws"
)

func main(){

	host:="http://localhost/testX"

	uri:="Zodos"
	tempHost:=host[len(host)-1:]
	temoUri:=uri[:1]

	fmt.Printf("host %v | tempUrl %v",tempHost,temoUri)
}



//Exemplo de utilizaçao dos pacotes para fazer um request simples com o metodo http Get
func GetExample(){

	c := client_ws.RequestTemplate{Url: "https://jsonplaceholder.typicode.com/"}
	c.AddHeader(client_ws.RequestHeader{Name: "user-id",Value:"123456"})
	var response = Book{}

	c.Get("todos/1").Body(&response)

	fmt.Println("Reposta service ",response)

}


func GetList()[]ProdutoDTO{
	var client = client_ws.RequestTemplate{Url: "http://localhost:3636/api/"}

	response :=make([]ProdutoDTO,0)

	if resp:=client.Get("produtos");resp.Err==nil{
		  resp.Body(&response)
		return response
	} else {
		fmt.Println("Erro a busar produtos",resp.Err)
		return nil
	}
	return nil
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

type  ProdutoDTO struct {
	Sku  int   `json:"sku"`
	Name  int   `json:"name"`
	Type  int   `json:"type"`
	Price float32 `json:"Price"`
	Shipping float32 `json:"Shipping"`
	ImageUrl string  `json:"image"`
}

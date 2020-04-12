package main

import (
	"fmt"
	"github.com/Andreson/go-portifolio/vitrine-produtos/controller/produto"
	"net/http"
)

func main(){

	//c := client_ws.RequestTemplate{Url: "https://jsonplaceholder.typicode.com/"}
	//c.AddHeader(client_ws.RequestHeader{Name: "user-id",Value:"123456"})
	//var response = Book{}
	//
	//c.Get("todos/1").Body(&response)
	//
	//fmt.Println("Resposta WS ", response)

	http.HandleFunc("/produtos", produto.Home);


	http.HandleFunc("/detalhes", produto.Detail);

	http.HandleFunc("/finalizarcompra", produto.FinalizarCompra)

	fmt.Println("Servidor iniciado na porta :9099")
	http.ListenAndServe(":9099",nil)

}


type Book struct {
	Id   int    `json:"userId"`
	Nome string `json:"title"`
}
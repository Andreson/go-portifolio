package service

import (
	"fmt"
	client_ws "github.com/Andreson/go-portifolio/ws-rest-wrapper/client-ws"
	dao "github.com/Andreson/go-portifolio/persistence-wrapper/dao"
	"strconv"
)

var client = client_ws.RequestTemplate{Url: "http://localhost:3636/api/"}

func Getall()[]ProdutoDTO {
	var response = make([]ProdutoDTO,0)
	if resp:=client.Get("produtos");resp.Err==nil{
		resp.Body(&response)
		return response
	} else {
		fmt.Println("Erro a busar produtos",resp.Err)
		return nil
	}
}


func FindOne(sku int)ProdutoDTO {
	var response =make([]ProdutoDTO,0)
	if resp:=client.Get("produtos?sku="+strconv.Itoa(sku));resp.Err==nil{
		resp.Body(&response)
		return response[0]
	} else {
		fmt.Println("Erro a busar produtos",resp.Err)
		return ProdutoDTO{}
	}
}

func FinishPurchase(sku int)ProdutoDTO{
	  produto:=FindOne(sku)

	dao.Save(produto.parseToEntity())
	return produto
}

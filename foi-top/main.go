package main

import (
	aut_controller "github.com/Andreson/go-portifolio/foi-top/controller/autentication"
	coast_controller "github.com/Andreson/go-portifolio/foi-top/controller/coast"
	event_controller "github.com/Andreson/go-portifolio/foi-top/controller/event"
	user_controller "github.com/Andreson/go-portifolio/foi-top/controller/usuario"
	"log"
	"net/http"
	"os"
)

/**
Projeto para fim de aprendizado da linguagem golang
Libs utilizadas
  - jwt-go  - Geraçao de token jwt
  - gorm    - Framework de persistencia SQL ORM
API rest com a finalidade de persistir registros no banco de dados myql
Foram implementados recurso jwt-go de  geraçao de token JWT
e Http Filter handle com recursos nativos da linguagem

  #Env PARAM
   - HOST_PORT
   - PROFILE - UTILIZADO PARA DEFINIR SE A CONEXAO DO BANCO DE DADOS
SER LOCAL OU REMOTA
       - dev : conexao local
       - others: conexao remota
 */

func  main()  {

	event_controller.Init()
	coast_controller.Init()
	aut_controller.Init()
	user_controller.Init()

	var port=""
	if port =os.Getenv("HOST_PORT");port=="" {
		port="8282"
	}

	log.Printf("Servidor iniciado na porta :%s\n",port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
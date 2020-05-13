package main

import (
	"crypto/sha1"
	"encoding/base64"
	aut_controller "github.com/Andreson/go-portifolio/foi-top/controller/autentication"
	coast_controller "github.com/Andreson/go-portifolio/foi-top/controller/coast"
	event_controller "github.com/Andreson/go-portifolio/foi-top/controller/event"
	user_controller "github.com/Andreson/go-portifolio/foi-top/controller/usuario"
	notific_service "github.com/Andreson/go-portifolio/foi-top/service/pubsub"
	"log"
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

	sh :=sha1.New()
	sh.Write([]byte("login.Password"))
	pass :=sh.Sum(nil)

	log.Println("############### value of parsing sha1",base64.URLEncoding.EncodeToString(pass))

	log.Printf("Servidor iniciado na porta :%s\n",port)

	notific_service.SendMesage("send-invite-event", Test{Nome:"Bruce", Idade: 54})


	//log.Fatal(http.ListenAndServe(":"+port, nil))
}

type Test struct {
	Nome string
	Idade int
}
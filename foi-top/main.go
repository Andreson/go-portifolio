package main

import (
	aut_controller "github.com/Andreson/go-portifolio/foi-top/controller/autentication"
	coast_controller "github.com/Andreson/go-portifolio/foi-top/controller/custos"
	event_controller "github.com/Andreson/go-portifolio/foi-top/controller/evento"
	"log"
	"net/http"
	"os"
)

func  main()  {



	event_controller.Init()
	coast_controller.Init()
	aut_controller.Init()


	var port=""
	if port =os.Getenv("HOST_PORT");port=="" {
		port="8282"
	}

	log.Printf("Servidor iniciado na porta :%s\n",port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
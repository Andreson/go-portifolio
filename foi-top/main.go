package main

import (
	event_controller "github.com/Andreson/go-portifolio/foi-top/controller/evento"
	"log"
	"net/http"
)

func  main()  {

	http.HandleFunc("/event", event_controller.Create)
	log.Print("Servidor iniciado na porta :8282")
	log.Fatal(http.ListenAndServe(":8282", nil))
}
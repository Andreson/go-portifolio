package client_ws

import (
	"net/http"
	"time"
)

var client = &http.Client{Timeout: time.Second * 30}

type Request struct {
	Url string
}



type ClientRest interface {
	Get(uri string, respoModel interface{} )
	Post(uri string, respoModel interface{} )
}
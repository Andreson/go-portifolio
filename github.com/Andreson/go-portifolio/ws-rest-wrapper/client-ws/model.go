package client_ws

import (
	"net/http"
	"time"
)

var client = &http.Client{Timeout: time.Second * 30}

type RequestTemplate struct {
	Url string
	Auth RequestHeader
	Headers []RequestHeader
}


type ResponseTemplate struct {
	BodyData []byte //array de bytes com a respota do serviço serializada
	err error
	StatusCode int

}

type RequestHeader struct {
	Name string
	Value string
}


type ClientRest interface {
	Get(uri string, respoModel interface{} )
	Post(uri string, respoModel interface{} )
}

func (rt ResponseTemplate) Body( resp interface{}){
	unmarshalReponse(rt.BodyData, &resp)
}

func (rt RequestTemplate) AddHeader(rh RequestHeader){
	rt.Headers =append(rt.Headers,rh)
}
package client_ws

import (
	"net/http"
	"time"
)

var client = &http.Client{Timeout: time.Second * 30}

type ClientRest interface {
	Get(uri string, respoModel interface{} )
	Post(uri string, respoModel interface{} )
}

type RequestTemplate struct {
	Url string
	Auth RequestHeader
	Headers []RequestHeader
	Body interface{}
	GetUrl func( ...string)

}


type ResponseTemplate struct {
	BodyData []byte //array de bytes com a respota do serviço serializada
	Err error
	StatusCode int

}

type RequestHeader struct {
	Name string
	Value string
}

func (rt ResponseTemplate) Body( resp interface{}) ResponseTemplate{
	unmarshalReponse(rt.BodyData, &resp)
	return rt
}

func (rt RequestTemplate) AddHeader(rh RequestHeader){
	rt.Headers =append(rt.Headers,rh)
}


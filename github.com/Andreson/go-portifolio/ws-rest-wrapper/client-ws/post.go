package client_ws

import (
	"bytes"
)

func (r RequestTemplate) Post(uri ...string ) (response ResponseTemplate){
	var urlRequest string
	if uri!=nil {
		urlRequest =r.Url+uri[0]
	}else {
		urlRequest=r.Url
	}

	var bodyRequest *bytes.Buffer
	if bodyRequest, response.Err =marshalRequest(r.Body); response.Err!=nil {
	return
	} else  {
	req :=createRequest(urlRequest,"POST", bodyRequest)
		byteResponse, err,httpReponse := call(req)
		response.StatusCode = httpReponse.StatusCode
		response.Err = err
		response.BodyData = byteResponse

	}
	return
}
package client_ws

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)


func unmarshalReponse(byteResponse []byte,respoModel interface{}) (error error){
	if error := json.Unmarshal(byteResponse, &respoModel); error == nil {
		return   nil
	} else {
		fmt.Println("Erro ao fazer Unmarshal: ",error)
		return   error
	}
}

func configHeaders(req *http.Request,auth RequestTemplate){
	var isContentType bool = false
	for _, v := range auth.Headers {
		if v.Name=="content-type" {isContentType=true}
		req.Header.Set(v.Name, v.Value)
	}
	if !isContentType {
		req.Header.Set("content-type","application-json; charset=utf-8")
	}

}

func configBasicAuth(req *http.Request,auth RequestTemplate){
		req.SetBasicAuth(auth.Auth.Name, auth.Auth.Value)
}


func call(req *http.Request) ([]byte, error,*http.Response) {
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer request", err)
		return nil, err,resp
	}

	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println("Erro ao fazer parse body response ", erro)
		return nil, erro,resp
	}
	return body, nil,resp
}

func createRequest(url string) *http.Request {

	if req, err := http.NewRequest("GET", url, nil); err == nil {
		return req
	} else {
		fmt.Println("Erro ao criar Request ", err)
		return nil
	}
}
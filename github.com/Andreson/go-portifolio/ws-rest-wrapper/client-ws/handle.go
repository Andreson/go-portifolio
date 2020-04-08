package client_ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r Request) Get(uri string, respoModel interface{} ) error{
	 urlRequest :=r.Url+uri
	req :=createRequest(urlRequest)
	if byteResponse, err := call(req); err == nil {

		if error := json.Unmarshal(byteResponse, &respoModel); error == nil {
			return   nil
		} else {
			fmt.Println("Erro ao fazer Unmarshal: ",error)
			return   error
		}
	} else {
		fmt.Println("Erro ao fazer request : ",err)
		return  errors.New("Erro ao fazer request ")
	}

}


func call(req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer request", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("HTTP Status code invalido " + string(resp.StatusCode))
	}

	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println("Erro ao fazer parse body response ", erro)
		return nil, erro
	}
	return body, nil
}

func createRequest(url string) *http.Request {

	if req, err := http.NewRequest("GET", url, nil); err == nil {
		return req
	} else {
		fmt.Println("Erro ao criar Request ", err)
		return nil
	}
}
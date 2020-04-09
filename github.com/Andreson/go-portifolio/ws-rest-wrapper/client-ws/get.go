package client_ws


func (r RequestTemplate) Get(uri string ) (response ResponseTemplate){
	urlRequest :=r.Url+uri
	req :=createRequest(urlRequest)
	byteResponse, err,httpReponse := call(req)
	response.StatusCode = httpReponse.StatusCode
	response.err = err
	response.BodyData = byteResponse

	return
}

package client_ws

func (r RequestTemplate) Get(uri string ) (response ResponseTemplate){
	urlRequest :=r.Url+uri

	req :=createRequest(urlRequest,"GET",nil)
	byteResponse, err,httpReponse := call(req)
	response.StatusCode = httpReponse.StatusCode
	response.Err = err
	response.BodyData = byteResponse


	return
}

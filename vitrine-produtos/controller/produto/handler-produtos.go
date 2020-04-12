package produto

import (
	service "github.com/Andreson/go-portifolio/vitrine-produtos/service/produto"
	"github.com/Andreson/go-portifolio/vitrine-produtos/view/produto"
	"net/http"
	"strconv"
)

func Home(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprint(writer,"Listagem de clientes ")
	produtos:=service.Getall()

	produto.ProdutoTemplate.ExecuteTemplate(writer,"listagem.html", &produtos)

}

func Detail(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprint(writer,"Listagem de clientes ")
	var resp service.ProdutoDTO
	 if sku, err :=strconv.Atoi( request.URL.Query().Get("sku"));err==nil {
		 resp = service.FindOne(sku)
	 }
	produto.ProdutoDetailTemplate.ExecuteTemplate(writer,"detalhe.html", &resp)

}

func FinalizarCompra(writer http.ResponseWriter, request *http.Request)  {


	var resp service.ProdutoDTO
	if sku, err :=strconv.Atoi( request.URL.Query().Get("sku"));err==nil {
		resp = service.FinishPurchase(sku)
	}
	produto.FinalizarCompra.ExecuteTemplate(writer,"finalizar.html", &resp)

}
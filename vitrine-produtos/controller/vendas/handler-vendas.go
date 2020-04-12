package vendas

import (
	vendas_service "github.com/Andreson/go-portifolio/vitrine-produtos/service/vendas"
	"github.com/Andreson/go-portifolio/vitrine-produtos/view/vendas"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request){

	vendas.VendasListagemTemplate.ExecuteTemplate(w,"listagem.html",vendas_service.GetAll())

}
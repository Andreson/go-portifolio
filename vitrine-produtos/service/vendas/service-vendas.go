package vendas_service

import (
	"database/sql"
	pw_dao "github.com/Andreson/go-portifolio/persistence-wrapper/dao"
	service "github.com/Andreson/go-portifolio/vitrine-produtos/service/produto"
)

func GetAll()[]service.ProdutoEntity{

	var query ="select Sku,Price,Name from produtoentity"
		vendas := make([]service.ProdutoEntity,0);

	callback := func(stmt *sql.Rows ){
		var sku int
		var price float64
		var name string

		stmt.Scan(&sku,&price,&name)

		vendas = append(vendas,service.ProdutoEntity{Sku:sku,Price:price,Name:name})
	}

	pw_dao.Query(query,callback)

	return  vendas;

}
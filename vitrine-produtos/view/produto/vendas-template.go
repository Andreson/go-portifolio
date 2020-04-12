package produto

import "html/template"

var ProdutoTemplate =template.Must(template.ParseFiles("./view/produto/listagem.html"))

var ProdutoDetailTemplate =template.Must(template.ParseFiles("./view/produto/detalhe.html"))

var FinalizarCompra =template.Must(template.ParseFiles("./view/produto/finalizar.html"))
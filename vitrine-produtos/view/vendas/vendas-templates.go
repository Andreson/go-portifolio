package vendas

import "html/template"

var VendasListagemTemplate = template.Must(template.ParseFiles("./view/vendas/listagem.html"))
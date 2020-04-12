package service

type  ProdutoDTO struct {
	Sku  int   `json:"sku"`
	Name  string   `json:"name"`
	Type  string   `json:"type"`
	Price float64 `json:"price"`
	Description string `json:"description"`
	ImageUrl string  `json:"image"`
}


func (p ProdutoDTO) Mod (dividendo int, divisor int) bool {
	return dividendo % divisor ==0
}

type ProdutoEntity struct {
	Sku int
	Price float64
	Name string
}


func (p ProdutoDTO) parseToEntity()ProdutoEntity{
	return ProdutoEntity{Sku:p.Sku, Price:p.Price, Name: p.Name}
}


package event_entity

//struct que representa um idem de despesa do evento, pode ser cerveja, arroz, agua etc

type ItemDespesaEventoEntity struct {
	BaseModel
	Nome string
	Descricao string
	Valor float32
	Unidade int
	EventoRefer int

}




func (ItemDespesaEventoEntity) TableName()string {
	return "despesaEvento"
}

type UnidadeMedida int

const (
	Caixa UnidadeMedida = iota
	Kilo
	Litro
	Pacote
	Unidade
)

func (d UnidadeMedida) String()string{
	return [...] string{"Caixa","KG","Litro","Pacote","Unidade"}[d]
}
package event_entity

import "time"

type  EventoEntity struct {
	BaseModel
	Titulo string
	Endereco string
	Bairro string
	Data time.Time
	Status string
	UserID int
  	Despesas []ItemDespesaEventoEntity `gorm:"foreignkey:EventoRefer"`
}

func (EventoEntity) TableName() string {
	return "evento"
}
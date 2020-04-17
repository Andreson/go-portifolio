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
}

func (EventoEntity) TableName() string {
	return "profiles"
}
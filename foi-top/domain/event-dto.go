package domain

import (
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"time"
)

type EventDto struct {
	Id       int64
	Titulo   string
	Endereco string
	Bairro   string
	Data     string
	Status   EventoStatus
	Usuario  UserDTO
}
type EventoStatus int
const (
	Confirmado EventoStatus = iota
	Criado
	Cancelado
)
func (d EventoStatus) String() string {
	return [3]string{"Confirmado", "Criado", "Cancelado"}[d]
}
func (e EventDto) ToEntity() (event_entity.EventoEntity) {
	return event_entity.EventoEntity{
		BaseModel:event_entity.BaseModel{ID:e.Id} ,
		Titulo: e.Titulo,
		Endereco: e.Endereco,
		Bairro: e.Bairro,
		UserID: e.Usuario.Id ,
		Data: time.Now()}
}
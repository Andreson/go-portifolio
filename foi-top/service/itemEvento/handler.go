package itemevento_service

import (
	"github.com/Andreson/go-portifolio/foi-top/domain"
	despesa_dao "github.com/Andreson/go-portifolio/foi-top/persistence/dao/despesa"
	event_entity "github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
)

func Cadastrar( dto domain.ItemDespesaEventoDTO) {
	e :=ToEntity(dto)

	despesa_dao.Save(e)
}

func ListByEvent(eventId int )[]domain.ItemDespesaEventoDTO{
	var entityQuery []event_entity.ItemDespesaEventoEntity
	despesa_dao.ListByEvent(eventId, &entityQuery)
	response := make( []domain.ItemDespesaEventoDTO, len(entityQuery))

	for i, v :=  range entityQuery {
		response[i]= ToDto(v)
	}

	return response
}

func ToDto(item event_entity.ItemDespesaEventoEntity ) domain.ItemDespesaEventoDTO {
	return domain.ItemDespesaEventoDTO{
		EventoId:item.ID,
		Valor:item.Valor,
		Unidade:item.Unidade,
		Descricao:item.Descricao,
		Nome:item.Nome}

}


func ToEntity(item domain.ItemDespesaEventoDTO) event_entity.ItemDespesaEventoEntity {

	return  event_entity.ItemDespesaEventoEntity {
		Nome:item.Nome,
		Descricao:item.Descricao,
		EventoRefer:  item.EventoId  ,
		Unidade:item.Unidade,
		Valor:item.Valor,
	}
}
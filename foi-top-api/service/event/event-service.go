package event_service

import (
	"errors"
	"github.com/Andreson/go-portifolio/foi-top/domain"
	evento_dao "github.com/Andreson/go-portifolio/foi-top/persistence/dao/evento"
	"github.com/Andreson/go-portifolio/foi-top/persistence/entitys"
	"log"
)

func validatedEvent(e domain.EventDto) error{
	var validateErrors = make([]error, 5)

	if e.Titulo=="" {
		validateErrors = append(validateErrors, errors.New("Titulo do evendo nao informado") )
	}

	return  nil
}

func Save( dto domain.EventDto) {
	e :=dto.ToEntity()
	log.Println("Cadastrando evento ",e)

	evento_dao.Save(e)
}

func FindById(dto domain.EventDto) event_entity.EventoEntity {

	return evento_dao.FindById(dto.ToEntity())
}
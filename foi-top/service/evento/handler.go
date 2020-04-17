package event_service

import (
	"errors"
	event_model "github.com/Andreson/go-portifolio/foi-top/model/evento"
	"github.com/Andreson/go-portifolio/foi-top/persistence"
	"log"
)

func validarEvento(e event_model.EventoDto) error{
	var validateErrors = make([]error, 5)

	if e.Titulo=="" {
		validateErrors = append(validateErrors, errors.New("Titulo do evendo nao informado") )
	}

	return  nil
}

func Cadastrar( dto event_model.EventoDto) {
	e :=dto.ToEntity();
	log.Println("Cadastrando evento ",e)

	persistence.Save(e)
}
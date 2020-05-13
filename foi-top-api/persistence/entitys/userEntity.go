package event_entity

type UsuarioEntity struct {
	BaseModel
	Nome string
	Email string
	Fone string
	Endereco string
}


func (UsuarioEntity) TableName()string {
	return "usuario"
}
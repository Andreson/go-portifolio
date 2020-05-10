package event_entity



type UserByEventEntity struct {
	BaseModel
	User UsuarioEntity `gorm:"association_foreignkey:id"`
	Evento EventoEntity `gorm:"association_foreignkey:id"`
	UserId int64
	EventoId int64

}

func (UserByEventEntity) TableName()string {
	return "usuario_evento"
}
package domain



type UserDTO struct {
	Nome     string
	Email    string
	Fone     string
	Endereco string
	Id       int64
	Login    LoginDTO
}
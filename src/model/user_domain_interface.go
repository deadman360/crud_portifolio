package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)

	GetBSONValue() (interface{}, error)

	EncryptPassword()
}

func (ud *userDomain) SetID(id string) {
	ud.Id = id
}

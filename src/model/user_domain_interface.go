package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetId() string

	SetID(string)

	EncryptPassword()
}

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

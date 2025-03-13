package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/deadman360/crud_portifolio/src/configuration/logger"
	"go.uber.org/zap"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8

	SetID(string)

	GetJSONValue() ([]byte, error)

	EncryptPassword()
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	Id       string
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age"`
}

var (
	journey = zap.String("journey", "UserDomain")
)

func (ud *userDomain) GetJSONValue() ([]byte, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		logger.Error("Error parsing userDomain to JSON", err, journey)
		return nil, err
	}
	return b, nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) SetID(id string) {
	ud.Id = id
}

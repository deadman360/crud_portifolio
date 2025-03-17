package converter

import (
	"github.com/deadman360/crud_portifolio/src/controller/repository/entity"
	"github.com/deadman360/crud_portifolio/src/model"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}

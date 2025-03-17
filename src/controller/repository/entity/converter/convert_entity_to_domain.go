package converter

import (
	"github.com/deadman360/crud_portifolio/src/controller/repository/entity"
	"github.com/deadman360/crud_portifolio/src/model"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)
	domain.SetID(entity.ID)
	return domain
}

package view

import (
	"github.com/deadman360/crud_portifolio/src/controller/routes/model/response"
	"github.com/deadman360/crud_portifolio/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

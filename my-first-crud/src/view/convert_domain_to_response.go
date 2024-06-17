package view

import (
	"github.com/arturbaccarin/go-my-first-crud/src/controller/model/response"
	"github.com/arturbaccarin/go-my-first-crud/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		Id:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}

package view

import (
	"github.com/dan-santos/go-user/src/controller/dto"
	"github.com/dan-santos/go-user/src/model"
)

func UserPresenter(
	userDomain model.UserDomainInterface,
) dto.UserResponse {
	return dto.UserResponse{
		ID: userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(),
	}
}
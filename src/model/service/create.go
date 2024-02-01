package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
)

func (user *userDomainService) CreateService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *resterrors.RestErr) {


	userAlreadyExists, _ := user.FindByEmailService(userDomain.GetEmail())
	if userAlreadyExists != nil {
		return nil, resterrors.NewBadRequestError("User with same email already exists")
	}

	userDomain.EncryptPassword()

	userDomainRepository, err := user.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}
	
	return userDomainRepository, nil
}
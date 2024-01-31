package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
)

func (user *userDomainService) FindByIdService(id string) (
	model.UserDomainInterface, *resterrors.RestErr,
) {
	return user.userRepository.FindUserById(id)
}

func (user *userDomainService) FindByEmailService(email string) (
	model.UserDomainInterface, *resterrors.RestErr,
) {
	
	return user.userRepository.FindUserByEmail(email)
}
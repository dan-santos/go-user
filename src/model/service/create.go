package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
)

func (user *userDomainService) CreateService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *resterrors.RestErr) {
	userDomain.EncryptPassword()

	userDomainRepository, err := user.userRepository.CreateUser(userDomain)
	if err != nil {
		return nil, err
	}
	
	return userDomainRepository, nil
}
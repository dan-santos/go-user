package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
)

func (user *userDomainService) LoginService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *resterrors.RestErr) {

	userDomain.EncryptPassword()

	userDB, err := user.findByEmailAndPasswordService(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, err
	}

	
	return userDB, nil
}
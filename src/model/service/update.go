package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
)

func (user *userDomainService) UpdateService(
		id string, 
		userDomain model.UserDomainInterface,
	) (*resterrors.RestErr) {
	err := user.userRepository.UpdateUser(id, userDomain)
	if err != nil {
		return err
	}
	return nil
}
package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
)

func (user *userDomainService) DeleteService(id string) *resterrors.RestErr {
	err := user.userRepository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
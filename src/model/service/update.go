package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
)

func (user *userDomainService) Update(
		userId string, 
		userDomain model.UserDomainInterface,
	) (*resterrors.RestErr) {
	
	return nil
}
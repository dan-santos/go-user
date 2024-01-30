package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
)

func (user *userDomainService) Find(string) (
	*model.UserDomainInterface, *resterrors.RestErr,
) {
	
	return nil, nil
}
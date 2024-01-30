package service

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
	"github.com/dan-santos/go-user/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository: userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	Create(model.UserDomainInterface) (model.UserDomainInterface, *resterrors.RestErr)
	Update(string, model.UserDomainInterface) *resterrors.RestErr
	Find(string) (*model.UserDomainInterface, *resterrors.RestErr)
	Delete(string) *resterrors.RestErr
}
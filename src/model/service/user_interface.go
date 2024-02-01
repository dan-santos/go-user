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
	CreateService(model.UserDomainInterface) (model.UserDomainInterface, *resterrors.RestErr)
	UpdateService(string, model.UserDomainInterface) *resterrors.RestErr
	FindByIdService(id string) (model.UserDomainInterface, *resterrors.RestErr)
	FindByEmailService(email string) (model.UserDomainInterface, *resterrors.RestErr)
	DeleteService(id string) *resterrors.RestErr
}
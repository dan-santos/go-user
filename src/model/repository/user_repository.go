package repository

import (
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_NAME = "COLLECTION_NAME"

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *resterrors.RestErr)
	
	FindUserById(
		id string,
	) (model.UserDomainInterface, *resterrors.RestErr)

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *resterrors.RestErr)

	FindUserByEmailAndPassword(
		email, password string,
	) (model.UserDomainInterface, *resterrors.RestErr)

	UpdateUser(
		id string,
		userDomain model.UserDomainInterface,
	) *resterrors.RestErr

	DeleteUser(
		id string,
	) *resterrors.RestErr
}
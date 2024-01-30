package repository

import (
	"context"
	"os"

	"github.com/dan-santos/go-user/src/configs/logger"
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
	"github.com/dan-santos/go-user/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var COLLECTION_NAME = os.Getenv("COLLECTION_NAME")

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *resterrors.RestErr) {
	collection := ur.databaseConnection.Collection(COLLECTION_NAME)

	document := entity.ToDatabase(userDomain)
	result, err := collection.InsertOne(context.Background(), document)
	if err != nil {
		logger.Error("Error on trying to insert user", err)
		return nil, resterrors.NewInternalServerError(err.Error())
	}
	
	document.ID = result.InsertedID.(primitive.ObjectID)
	
	return entity.ToDomain(*document), nil
}
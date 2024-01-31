package repository

import (
	"context"
	"os"

	"github.com/dan-santos/go-user/src/configs/logger"
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
	"github.com/dan-santos/go-user/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) UpdateUser(
	userId string,
	userDomain model.UserDomainInterface,
) *resterrors.RestErr {
	collection_name := os.Getenv(COLLECTION_NAME)
	collection := ur.databaseConnection.Collection(collection_name)

	id, _ := primitive.ObjectIDFromHex(userId)
	document := entity.ToDatabase(userDomain)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: document}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error on trying to update user", err)
		return resterrors.NewInternalServerError(err.Error())
	}
	
	return nil
}
package repository

import (
	"context"
	"os"

	"github.com/dan-santos/go-user/src/configs/logger"
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) DeleteUser(
	userId string,
) *resterrors.RestErr {
	collection_name := os.Getenv(COLLECTION_NAME)
	collection := ur.databaseConnection.Collection(collection_name)

	id, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: id}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error on trying to delete user", err)
		return resterrors.NewInternalServerError(err.Error())
	}
	
	return nil
}
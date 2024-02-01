package repository

import (
	"context"
	"fmt"
	"os"

	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/model"
	"github.com/dan-santos/go-user/src/model/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *resterrors.RestErr) {
	collection_name := os.Getenv(COLLECTION_NAME)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email,
			)
			return nil, resterrors.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		return nil, resterrors.NewInternalServerError(errorMessage)
	}
	return entity.ToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email, password string,
) (model.UserDomainInterface, *resterrors.RestErr) {
	collection_name := os.Getenv(COLLECTION_NAME)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}, {Key: "password", Value: password}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("Wrong email or password")
			return nil, resterrors.NewUnauthorizedError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		return nil, resterrors.NewInternalServerError(errorMessage)
	}
	return entity.ToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(
	id string,
) (model.UserDomainInterface, *resterrors.RestErr) {
	collection_name := os.Getenv(COLLECTION_NAME)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this id: %s", id,
			)
			return nil, resterrors.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		return nil, resterrors.NewInternalServerError(errorMessage)
	}
	return entity.ToDomain(*userEntity), nil
}
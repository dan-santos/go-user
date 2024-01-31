package controller

import (
	"net/http"

	"github.com/dan-santos/go-user/src/configs/logger"
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/configs/validation"
	"github.com/dan-santos/go-user/src/controller/dto"
	"github.com/dan-santos/go-user/src/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) Update(c *gin.Context) {
	var userRequest dto.UserUpdateRequest
	userId := c.Param("userId")
	
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("path", "updateUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := resterrors.NewBadRequestError("UserID is not a valid id")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateService(userId, domain); if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully")
	c.Status(http.StatusOK)
}
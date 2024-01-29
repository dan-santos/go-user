package controller

import (
	"net/http"

	"github.com/dan-santos/go-user/src/configs/logger"
	"github.com/dan-santos/go-user/src/configs/validation"
	"github.com/dan-santos/go-user/src/controller/dto"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	var userRequest dto.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("path", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	response := dto.UserResponse{
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	logger.Info("User created successfully")
	c.JSON(http.StatusOK, response)
}
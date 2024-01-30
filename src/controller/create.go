package controller

import (
	"net/http"

	"github.com/dan-santos/go-user/src/configs/logger"
	"github.com/dan-santos/go-user/src/configs/validation"
	"github.com/dan-santos/go-user/src/controller/dto"
	"github.com/dan-santos/go-user/src/model"
	"github.com/dan-santos/go-user/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) Create(c *gin.Context) {
	var userRequest dto.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("path", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	data, err := uc.service.Create(domain); if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully")
	c.JSON(http.StatusCreated, view.UserPresenter(data))
}
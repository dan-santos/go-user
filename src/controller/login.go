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

func (uc *userControllerInterface) Login(c *gin.Context) {
	var userRequest dto.UserLoginRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("path", "LoginUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	data, token, err := uc.service.LoginService(domain); if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User logged successfully")
	c.Header("Authorization", token)
	c.JSON(http.StatusCreated, view.UserPresenter(data))
}
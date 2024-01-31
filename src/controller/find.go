package controller

import (
	"net/http"
	"net/mail"

	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/dan-santos/go-user/src/view"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := resterrors.NewBadRequestError("UserID is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindByIdService(userId)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.UserPresenter(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		errorMessage := resterrors.NewBadRequestError("User mail is not a valid mail")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindByEmailService(userEmail)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.UserPresenter(userDomain))
}
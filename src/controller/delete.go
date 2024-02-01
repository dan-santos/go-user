package controller

import (
	"net/http"

	"github.com/dan-santos/go-user/src/configs/logger"
	resterrors "github.com/dan-santos/go-user/src/configs/rest_errors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) Delete(c *gin.Context) {
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := resterrors.NewBadRequestError("UserID is not a valid id")
		c.JSON(errRest.Code, errRest)
	}
	err := uc.service.DeleteService(userId); if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User deleted successfully")
	c.Status(http.StatusOK)
}
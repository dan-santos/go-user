package controller

import (
	"fmt"

	"github.com/dan-santos/go-user/src/configs/validation"
	"github.com/dan-santos/go-user/src/controller/dto"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var userRequest dto.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
	}

	fmt.Println(userRequest)
}
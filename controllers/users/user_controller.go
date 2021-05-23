package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nononsensecode/bookstore_users-api/domain/users"
	"github.com/nononsensecode/bookstore_users-api/services"
	"github.com/nononsensecode/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	savedUser, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, savedUser)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewBadRequestError("user id should be a number")
		c.JSON(restErr.Status, restErr)
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
	}

	c.JSON(http.StatusOK, user)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented")
}

package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ramski29/bookstore_users-api/domain/user"
	"github.com/ramski29/bookstore_users-api/services"
	"github.com/ramski29/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user user.User
	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemented Me!")
}

//
//func FindUser(c *gin.Context)  {
//	c.String(http.StatusNotImplemented, "Implemented Me!")
//}

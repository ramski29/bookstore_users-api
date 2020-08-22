package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemented Me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implemented Me!")
}

//
//func FindUser(c *gin.Context)  {
//	c.String(http.StatusNotImplemented, "Implemented Me!")
//}

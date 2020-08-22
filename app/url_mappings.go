package app

import (
	"github.com/ramski29/bookstore_users-api/controllers/ping"
	"github.com/ramski29/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:id_user", users.GetUser)
	//router.GET("/users/search", controllers.FindUser)
	router.POST("/create", users.CreateUser)
}

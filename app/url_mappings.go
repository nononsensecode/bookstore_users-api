package app

import (
	"github.com/nononsensecode/bookstore_users-api/controllers/ping"
	"github.com/nononsensecode/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/users/search", users.FindUser)
	router.POST("/users", users.Create)
	router.PUT("/users", users.Update)
	router.PATCH("/users", users.Update)
}

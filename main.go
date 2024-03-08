package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sebmaz93/book_my_event/db"
	"github.com/sebmaz93/book_my_event/routes"
)

func main() {
	db.Init()
	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":777")
	if err != nil {
		return
	}
}

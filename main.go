package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jxnhiro/event-api/go/db"
	"github.com/jxnhiro/event-api/go/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
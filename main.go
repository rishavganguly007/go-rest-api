package main

import (
	
	"example.com/go-rest-api/db"
	"example.com/go-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default() // configures a http server
	
	routes.RegisterRoutes(server)
	
	server.Run(":8080") //localhost port

}


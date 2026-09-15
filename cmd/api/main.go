package main

import (
	"book.api/train4/internal/database"
	"book.api/train4/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	r := gin.Default()
	routes.Register(r)

	r.Run(":8080")
}

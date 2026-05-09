package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//cfg := config.LoadConfig()

	router.Run(":8080")

}

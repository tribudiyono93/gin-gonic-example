package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	gin.DisableConsoleColor()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Fatal(r.Run(":8089"))
}

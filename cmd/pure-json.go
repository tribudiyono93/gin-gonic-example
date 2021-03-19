package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"html": "<b>hello world</b>",
		})
	})
	
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>hello world</b>",
		})
	})

	log.Panic(r.Run(":8080"))
}

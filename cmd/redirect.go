package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.com")
	})

	r.GET("/foo", func(c *gin.Context) {
		c.Request.URL.Path = "/bar"
		r.HandleContext(c)
	})

	r.GET("/bar", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	
	log.Panic(r.Run(":8080"))
}

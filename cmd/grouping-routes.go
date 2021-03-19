package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/something", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK")
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/something", func(c *gin.Context) {
			c.JSON(http.StatusOK, "OK")
		})
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run", err)
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding="required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.POST("/loginXML", func(c *gin.Context) {
		var xml Login
		if err := c.ShouldBindXML(&xml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if xml.User != "manu" || xml.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if form.User != "manu" || form.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})

	})

	log.Panic(r.Run(":8080"))
}

/*
curl -v -X POST \
  http://127.0.0.1:8080/loginJSON \
  -H 'content-type: application/json' \
  -d '{ "user": "manu" }'
 */

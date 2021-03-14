package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person2 struct {
	ID string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/:name/:id", func(c *gin.Context) {
		var person Person2
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "uuid": person.ID})
	})

	log.Fatal(r.Run(":8088"))
}

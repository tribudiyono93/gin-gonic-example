package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person3 struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	r := gin.Default()
	r.Any("/testing", startPage2)
	r.Run(":8080")
}

func startPage2(c *gin.Context) {
	var person Person3
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(http.StatusOK, "Success")
}

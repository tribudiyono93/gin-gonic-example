package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct{
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStructPointer,
		"b": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/get-b", GetDataB)
	r.GET("/get-c", GetDataC)
	r.GET("/get-d", GetDataD)

	log.Fatal(r.Run(":8089"))
}
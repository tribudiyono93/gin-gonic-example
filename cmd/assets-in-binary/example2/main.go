package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	var f embed.FS

	r := gin.Default()
	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl", "templates/foo/*.tmpl"))
	r.SetHTMLTemplate(templ)

	r.StaticFS("/public", http.FS(f))
	
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bar.tmpl", gin.H{
			"title": "Foo website",
		})
	})

	r.GET("favicon.ico", func(c *gin.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		c.Data(http.StatusOK, "image/x-icon", file)
	})

	r.Run(":8080")
}
package main

import (
	"github.com/gin-gonic/gin"
	"uso.link/internal/handler"
)

func main() {

}
func DoTestHtml() {
	g := gin.New()
	g.LoadHTMLGlob("vcp/*")
	g.Static("/static", "public/static/")
	g.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/index.html")
	})
	g.GET("/:name.html", handler.RenderPage)
	g.Run(":2222")
}

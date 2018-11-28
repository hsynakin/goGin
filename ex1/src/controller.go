package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Kontrol() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)

	})
	return r
}

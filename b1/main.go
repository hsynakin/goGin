package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hsynakin/goGin/b1/apiroots"
)

func main() {
	app := gin.Default()

	api := app.Group("/api")

	apiroots.Einvoiceservices(api)

	app.Run(fmt.Sprintf(":%d", 1455))
}

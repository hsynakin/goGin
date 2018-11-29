package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/thename", func(c *gin.Context) {
		firstname := c.DefaultQuery("Firtname", "john")
		lastname := c.DefaultQuery("lastname", "foo")
		c.JSON(200, gin.H{
			"firstname": firstname,
			"lastname":  lastname,
		})
	})
	r.Run(":8000")
}

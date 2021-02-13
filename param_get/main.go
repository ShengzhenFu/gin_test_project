package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.Query("last_name")
		level := c.DefaultQuery("level", "rookie")
		c.String(http.StatusOK, "%s, %s, %s", firstName, lastName, level)
	})
	r.Run(":8612")
}

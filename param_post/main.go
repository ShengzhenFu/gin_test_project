package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.PostForm("last_name")
		level := c.DefaultPostForm("level", "Rookie")
		c.String(http.StatusOK, "%s, %s, %s", firstName, lastName, level)
	})
	r.Run(":8612")
}

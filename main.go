package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/books", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Books",
			"time": time.Now(),
		})
	})

	r.Run()
}

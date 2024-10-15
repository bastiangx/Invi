package main

import (
	"fmt"
	"net/http"
  "os"

	"github.com/gin-gonic/gin"
  "github.com/"
)

func main() {
	// hello world gin
	fmt.Println("Hello World")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}

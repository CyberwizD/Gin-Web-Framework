package main

import (
	"github.com/gin-gonic/gin"
)

// This is a simple example of a Gin web server that responds to a GET request at the "/ping" endpoint.
// When a request is made to this endpoint, it responds with a JSON object containing the message "pong".
// The server listens on port 8080 by default.

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

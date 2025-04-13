package main

import (
	"github.com/gin-gonic/gin"
)

// Redirects is a middleware that redirects the request to a different URL based on the request path.

func v1Handler(c *gin.Context) {
	// Redirect to v2Handler
	c.Redirect(302, "/api/v2")
}

func v2Handler(c *gin.Context) {
	// Handle the request for v2
	c.JSON(200, gin.H{
		"message": "Welcome to v2!",
	})
}

func main() {
	r := gin.Default()

	r.GET("/api/v1", v1Handler) // Redirect to v2Handler
	r.GET("/api/v2", v2Handler) // Handle the request for v2

	r.Run(":8080")
}

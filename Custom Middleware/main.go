package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// Before request
		log.Println("Before request")
		log.Println("Request URL:", c.Request.URL.Path)
		c.Next()

		// After request
		latency := time.Since(t)
		log.Println("After request")
		log.Println("Request Method:", c.Request.Method)
		log.Print(latency)

		// Access the status from the context
		// c.Writer.Status() will return the HTTP status code of the response sent to the client

		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	router := gin.New()

	router.Use(Logger())

	router.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// Prints "12345"
		log.Println(example)
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")
}

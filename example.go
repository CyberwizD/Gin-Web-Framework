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

	r.POST("/register", func(c *gin.Context) {
		var json struct {
			Id       int64  `json:"id"`
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		if json.Username == "admin" && json.Password == "password" {
			json.Id = 1                // Assign a admin ID
			json.Username = "admin"    // Assign a admin username
			json.Password = "password" // Assign a admin password

			c.JSON(200, gin.H{"status": "registered"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	})

	r.GET("/admin/:id", func(c *gin.Context) {
		id := c.Param("id") // Get the user ID from the URL parameter
		c.JSON(200, gin.H{
			"admin_id": id, // Respond with the user ID
			"message":  "admin found",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

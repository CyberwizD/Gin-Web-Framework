package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type formA struct {
	Name string `json:"name" xml:"name" binding:"required"`
}

type formB struct {
	Password string `json:"password" xml:"password" binding:"required"`
}

// c.ShouldBind consumes c.Request.Body and it cannot be reused, so c.ShouldBindBodyWith is used to bind the body into different structs.
// c.ShouldBindBodyWith is a low-level function that allows you to bind the request body into different structs without consuming it.
// It is useful when you want to bind the request body into different structs based on the content type or other conditions.

func FormHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	// Bind JSON to objA
	if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
		c.JSON(200, gin.H{
			"message": "Form A",
			"name":    objA.Name,
		})
	} else {
		c.JSON(400, gin.H{
			"error": errA.Error(),
		})
	}

	// Bind XML to objB
	if errB := c.ShouldBindBodyWith(&objB, binding.XML); errB == nil {
		c.JSON(200, gin.H{
			"message":  "Form B",
			"password": objB.Password,
		})
	} else {
		c.JSON(400, gin.H{
			"error": errB.Error(),
		})
	}
}

// c.ShouldBindBodyWith stores body into the context before binding.
// This has a slight impact to performance, so you should not use this method if you are enough to call binding at once.

// This feature is only needed for some formats — JSON, XML, MsgPack, ProtoBuf.
// For other formats, Query, Form, FormPost, FormMultipart, can be called by c.ShouldBind() multiple times without any damage to performance

func main() {
	r := gin.Default()

	r.POST("/form", FormHandler)

	r.Run(":8080")
}

package router

import (
	gin "github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/", func(c *gin.Context) {
		r.Static("/", "../front")
	})
	r.Run(":8080")
}

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//we can also gin.New(). Default() use logger, and recover
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8083")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MyLogger() gin.HandlerFunc {

	return func(c *gin.Context) {
		t := time.Now()
		//we can extract the jwt info, and set it here
		c.Set("example", "123455")
		//the chain continue
		c.Next()

		end := time.Since(t)
		fmt.Println(end)
		fmt.Println(c.Writer.Status())

	}
}

func main() {

	router := gin.Default()
	//router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	router.Use(MyLogger())

	auth := router.Group("/auth")
	{
		auth.Use(MyAuth())
		auth.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "pong",
			})
		})

	}

	router.Run(":8083")

}

func MyAuth() gin.HandlerFunc {
	return AuthRequired
}

func AuthRequired(c *gin.Context) {
	fmt.Println("AuthRequired")
}

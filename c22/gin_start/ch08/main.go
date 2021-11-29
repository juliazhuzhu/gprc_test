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
	router.Use(TokenRequired())

	auth := router.Group("/auth")
	{
		auth.Use(DummyFunction())
		auth.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "pong",
			})
		})

	}

	router.Run(":8083")

}

func DummyFunction() gin.HandlerFunc {
	return Dummy
}

func Dummy(c *gin.Context) {
	fmt.Println("Dummy")
}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			}
		}

		if token != "bobby" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "unauthed",
			})

			//return, CAN NOT be used here. c.Abort() instead!!!
			c.Abort()
		}

		c.Next()
	}
}

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/welcome", welcome)
	router.POST("/form_post", formPost)

	router.Run(":8083")
}

func formPost(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anon")
	id := c.Query("id")
	page := c.Query("page")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick,
		"id":      id,
		"page":    page,
	})
}

func welcome(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "xiaoye")
	lastName := c.DefaultQuery("lastname", "zhu")
	c.JSON(http.StatusOK, gin.H{
		"first_name": firstName,
		"last_name":  lastName,
	})

}

package main

import (
	"github.com/gin-gonic/gin"
	"hexmeet.com/grpctest/c22/gin_start/ch06/proto"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/moreJSON", moreJSON)
	router.GET("/someProtoBuf", returnProtoc)
	router.GET("/json", func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H {
			"html":"<b>Hello, world!</b>",
		})
	})
	router.GET("/purejson", func(c *gin.Context){
		c.PureJSON(http.StatusOK,gin.H {
			"html":"<b>Hello, world!</b>",
		})
	})

	router.Run(":8083")
}

func returnProtoc(c *gin.Context) {
	course := []string { "python", "golang", "micro"}
	user := &proto.Teacher{
		Name :"xiaoye",
		Course:course,
	}

	c.ProtoBuf(http.StatusOK,user)
}

func moreJSON(c *gin.Context) {

	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	
	msg.Name = "xiaoye"
	msg.Message = "this is a test"
	msg.Number = 20

	c.JSON(http.StatusOK, msg)
}

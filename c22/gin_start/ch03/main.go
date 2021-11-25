package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//we can also gin.New(). Default() use logger, and recover
	//路由分组
	goodsGroup := r.Group("/goods")
	{
		goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/add", creatGood)
		//变量作为url
		goodsGroup.GET(":id/:action", goodsDetail)
		//goodsGroup.GET(":id/*path",goodsPath)
	}

	r.Run(":8084")
}

func goodsPath(c *gin.Context) {
	id := c.Param("id")
	path := c.Param("path")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"path": path,
	})
}

func goodsDetail(c *gin.Context) {
	//获取url中的变量
	id := c.Param("id")
	action := c.Param("action")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})

}

func creatGood(context *gin.Context) {

}

func goodsList(context *gin.Context) {

}

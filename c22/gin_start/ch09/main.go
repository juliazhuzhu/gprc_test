package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	//加载包含子目录文件的程序，参考 22-15
	router := gin.Default()
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println(dir)
	//a relative directory
	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLGlob("templates/**/*") //包含子目录，和其中的文件
	//router.LoadHTMLFiles("templates/index.tmpl")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "imoc",
		})
	})

	router.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"name": "你好小朋友",
		})
	})

	router.Run(":8083")
}

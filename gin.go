package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"syscall"

	"github.com/gin-gonic/gin"
)

type person struct {
	id   string
	name string
	age  int
}

func openBrowser() {
	cmd := exec.Command(`cmd`, `/c`, `start`, `http://localhost:3000`)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/login", func(c *gin.Context) {
		name := c.Query("username")
		pwd := c.Query("password")
		fmt.Println(name, pwd)
		if name == "guozb" && pwd == "123456" {
			c.JSON(200, gin.H{
				"code":    200,
				"data":    true,
				"message": "✔️ 登录成功",
			})
		} else {
			c.JSON(200, gin.H{
				"code":    200,
				"data":    false,
				"message": "账户名或密码错误",
			})
		}

	})
	r.GET("/cat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "ooooooh",
			"data": person{
				id:   "I001",
				name: "曼尼克斯",
			},
		})
	})
	go openBrowser()
	r.Run(":3000")

}

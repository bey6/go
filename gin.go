package main

import (
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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    200,
		})
	})
	r.POST("/cat", func(c *gin.Context) {
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

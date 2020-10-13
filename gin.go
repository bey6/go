package main

import "github.com/gin-gonic/gin"

type person struct {
	id   string
	name string
	age  int
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
	r.Run()
}

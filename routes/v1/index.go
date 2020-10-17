package v1

import "github.com/gin-gonic/gin"

func index(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "不知火舞",
	})
}

package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

func Router(engine *gin.Engine) {
	engine.POST("/user/upsert", func(c *gin.Context) {
		var param User
		err := c.ShouldBindJSON(&param)
		if err != nil {
			c.JSON(200, gin.H{
				"code":    http.StatusBadRequest,
				"message": "参数错误",
			})
			return
		}
		c.JSON(http.StatusOK, "success")
	})
}

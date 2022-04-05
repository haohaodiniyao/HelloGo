package hello

import "github.com/gin-gonic/gin"

func Router(engine *gin.Engine) {
	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(200, map[string]interface{}{
			"message": "hello go",
		})
	})
}

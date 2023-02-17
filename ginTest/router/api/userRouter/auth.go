package userRouter

import "github.com/gin-gonic/gin"

func userAuthRouter(g *gin.RouterGroup) {
	g.GET("/login", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "l"})
	})
}

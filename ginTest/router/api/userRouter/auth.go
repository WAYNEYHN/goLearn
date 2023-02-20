package userRouter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func userAuthRouter(g *gin.RouterGroup) {
	//http://localhost:8000/api/user-auth/login/123/1231
	g.GET("/login/:username/:password", func(context *gin.Context) {
		account := context.Param("username")
		password := context.Param("password")
		context.JSON(http.StatusOK, gin.H{"message": account + "_" + password})
	})

	//http://localhost:8000/api/user-auth/login?username=123&password=123
	g.GET("/login", func(context *gin.Context) {

		username := context.Query("username")
		password := context.Query("password")
		fmt.Println(context.GetHeader("token"))
		//除了使用以上方法来
		if len(password) == 0 || len(username) == 0 {
			context.JSON(http.StatusOK,, gin.H{"message": "username and password shouldn't be null"})
		} else {
			context.JSON(http.StatusOK, gin.H{"message": username + "_" + password})
		}

	})
}

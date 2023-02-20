package router

import (
	"github.com/gin-gonic/gin"
	"main/middleware"
	"main/router/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.MyMiddlerWare())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	api.InitApi(r)
	return r
}

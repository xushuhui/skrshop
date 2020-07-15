package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skrshop-api/api"
	"skrshop-api/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandle())
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})

	router.POST("/login", api.Login)

	return router
}

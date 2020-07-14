package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skrshop-api/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandle())
	//router.Use(middleware.Logger(), gin.Recovery())
	// 添加 Get 请求路由
	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	authMiddle := middleware.AuthMiddleware()
	router.POST("/login", authMiddle.LoginHandler)
	//router.Use(authMiddle.MiddlewareFunc())

	//

	return router
}

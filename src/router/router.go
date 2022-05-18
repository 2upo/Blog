package router

import (
	"blog/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoutes(baseGroup *gin.RouterGroup) {

	authMiddleware := middleware.InitJWT()

	baseGroup.GET("healthcheck", healthcheck)
	baseGroup.POST("/login", authMiddleware.LoginHandler)
	auth := baseGroup.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}
}

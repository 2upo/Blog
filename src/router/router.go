package router

import (
	"blog/auth"
	"blog/user"

	"github.com/gin-gonic/gin"
)

func InitRoutes(baseGroup *gin.RouterGroup) {

	authMiddleware := auth.InitJWT()

	baseGroup.GET("healthcheck", healthcheck)
	userController := user.InitController()
	userController.InitRoutes(baseGroup)
	baseGroup.POST("/login", authMiddleware.LoginHandler)
	auth := baseGroup.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)

	}
}

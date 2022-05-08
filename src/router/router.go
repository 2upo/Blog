package router

import "github.com/gin-gonic/gin"

func InitRoutes(baseGroup *gin.RouterGroup) {
	baseGroup.GET("healthcheck", healthcheck)
}

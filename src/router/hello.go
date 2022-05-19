package router

import (
	"blog/user"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	u, _ := c.Get(user.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[user.IdentityKey],
		"userName": u.(*user.User).UserName,
		"text":     "Hello World.",
	})
}

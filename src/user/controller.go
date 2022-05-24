package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Service IUserService
}

// Constructor
func InitController() *UserController {
	var stateController UserController
	stateController.Service = InitUserService()

	return &stateController
}

func (controller *UserController) InitRoutes(baseGroup *gin.RouterGroup) {
	// Register routes
	userGroup := baseGroup.Group("/user")
	{
		userGroup.POST("/", controller.insertUser)
	}
}

func (controller *UserController) insertUser(ctx *gin.Context) {
	var user RegisterSchema
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.Service.Insert(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": user,
		})
	}
}

package router

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthcheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

type Healthcheck struct {
	Status string `json:"status"`
}

func (h *Healthcheck) String() string {
	serialized_resp, _ := json.Marshal(h)

	return string(serialized_resp)
}

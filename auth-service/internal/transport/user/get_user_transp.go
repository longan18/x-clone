package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUserTransp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transport list user",
	})
}

func GetUserTransp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transport get user",
	})
}
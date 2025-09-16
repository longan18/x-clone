package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUserTransp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transport update user",
	})
}
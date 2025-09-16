package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUserTransp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transport delete user",
	})
}
package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserTransp(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Transport create user",
	})
}

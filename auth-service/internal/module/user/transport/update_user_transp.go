package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUserTransp(db *gorm.DB) gin.HandlerFunc{
	return func (ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Transport update user",
		})
	}
}
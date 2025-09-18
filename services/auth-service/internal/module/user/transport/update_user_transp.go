package transport

import (
	"auth-service/internal/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUserTransp(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, common.ResponseData(map[string]interface{}{
			"message": "Transport update user",
		}))
	}
}

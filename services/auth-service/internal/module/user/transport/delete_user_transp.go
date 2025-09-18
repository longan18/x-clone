package transport

import (
	"auth-service/internal/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUserTransp(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, common.ResponseData(map[string]interface{}{
			"message": "Transport delete user",
		}))
	}
}

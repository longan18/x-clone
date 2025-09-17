package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUserTransp(db *gorm.DB) gin.HandlerFunc{
	return func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Transport delete user",
		})
	}
}
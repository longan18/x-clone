package transport

import (
	"auth-service/internal/module/user/business"
	storage "auth-service/internal/module/user/storage"
	"auth-service/internal/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListUserTransp(db *gorm.DB) gin.HandlerFunc{
	return func (ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Transport list user",
		})
	}
}

func GetUserTransp(db *gorm.DB) gin.HandlerFunc{
	return func (c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			
		}

		storage := storage.NewMySQLStorage(db)
		biz := business.NewGetUserBiz(storage)

		user, err := biz.GetUserById(id)
		if err != nil {

		}

		c.JSON(http.StatusOK, util.NewGetSuccess(map[string]interface{}{
			"param": id,
			"user": user,
		}))
	}
}
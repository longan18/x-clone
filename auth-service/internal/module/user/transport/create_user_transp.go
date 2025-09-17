package transport

import (
	"auth-service/internal/model"
	"auth-service/internal/module/user/business"
	storage "auth-service/internal/module/user/storage"
	"auth-service/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserTransp(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userReq model.UserRequest

		if err := c.ShouldBindJSON(&userReq); err != nil {
			c.JSON(http.StatusUnprocessableEntity, util.NewValidationError(err))
			return
		}

		storage := storage.NewMySQLStorage(db)
		biz := business.NewCreateUserBiz(storage)

		user, err := biz.CreateNewUser(c, userReq)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err)
			return
		}

		c.JSON(http.StatusOK, util.NewCreateSuccess(map[string]uint{
			"user_id": user.ID,
		}))
	}
}

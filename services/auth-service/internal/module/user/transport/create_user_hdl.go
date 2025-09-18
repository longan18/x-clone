package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/model"
	"auth-service/internal/module/user/business"
	storage "auth-service/internal/module/user/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userReq model.UserRequest

		if err := c.ShouldBindJSON(&userReq); err != nil {
			validationErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid request format")
			c.JSON(validationErr.StatusCode(), validationErr)
			return
		}

		storage := storage.NewMySQLStorage(db)
		biz := business.NewCreateUserBiz(storage)

		user, err := biz.CreateNewUser(c, &userReq)
		if err != nil {
			if defaultErr, ok := err.(*common.DefaultError); ok {
				c.JSON(defaultErr.StatusCode(), defaultErr)
			} else {
				internalErr := common.ErrInternalServerError.WithTrace(err)
				c.JSON(internalErr.StatusCode(), internalErr)
			}
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(map[string]int{
			"user_id": user.Id,
		}))
	}
}

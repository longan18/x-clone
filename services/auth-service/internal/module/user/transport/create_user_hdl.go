package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/module/user/business"
	"auth-service/internal/module/user/entity"
	ustr "auth-service/internal/module/user/storage"
	rstr "auth-service/internal/module/role/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUserHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userReq entity.UserRequest

		if err := c.ShouldBindJSON(&userReq); err != nil {
			validationErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid request format")
			c.JSON(validationErr.StatusCode(), validationErr)
			return
		}

		ustr := ustr.NewMySQLStorage(db)
		rstr := rstr.NewMySQLStorage(db)
		biz := business.NewCreateUserBiz(ustr, rstr)

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

package transport

import (
	"auth-service/internal/common"
	rstr "auth-service/internal/module/role/storage"
	"auth-service/internal/module/user/business"
	"auth-service/internal/module/user/entity"
	ustr "auth-service/internal/module/user/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUserHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid user ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		var userReq entity.UserUpdateRequest

		if err := c.ShouldBindJSON(&userReq); err != nil {
			validationErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid request format")
			c.JSON(validationErr.StatusCode(), validationErr)
			return
		}

		ustr := ustr.NewMySQLStorage(db)
		rstr := rstr.NewMySQLStorage(db)
		biz := business.NewUpdateUserBiz(ustr, rstr)

		user, err := biz.UpdateUser(c, id, &userReq)
		if err != nil {
			if defaultErr, ok := err.(*common.DefaultError); ok {
				c.JSON(defaultErr.StatusCode(), defaultErr)
			} else {
				internalErr := common.ErrInternalServerError.WithTrace(err)
				c.JSON(internalErr.StatusCode(), internalErr)
			}
			return
		}

		c.JSON(http.StatusOK, common.ResponseData(map[string]interface{}{
			"message": "User updated successfully",
			"user":    user,
		}))
	}
}

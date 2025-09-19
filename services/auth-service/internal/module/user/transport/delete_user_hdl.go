package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/module/user/business"
	ustr "auth-service/internal/module/user/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUserHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid user ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		ustr := ustr.NewMySQLStorage(db)
		biz := business.NewDeleteUserBiz(ustr)

		err = biz.DeleteUser(c, id)
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
			"message": "User deleted successfully",
		}))
	}
}

func SoftDeleteUserHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid user ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		ustr := ustr.NewMySQLStorage(db)
		biz := business.NewDeleteUserBiz(ustr)

		err = biz.SoftDeleteUser(c, id)
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
			"message": "User soft deleted successfully",
		}))
	}
}

package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/business"
	rstr "auth-service/internal/module/role/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteRoleHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid role ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		rstr := rstr.NewMySQLStorage(db)
		biz := business.NewDeleteRoleBiz(rstr)

		err = biz.DeleteRole(c, id)
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
			"message": "Role deleted successfully",
		}))
	}
}

func SoftDeleteRoleHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid role ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		rstr := rstr.NewMySQLStorage(db)
		biz := business.NewDeleteRoleBiz(rstr)

		err = biz.SoftDeleteRole(c, id)
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
			"message": "Role soft deleted successfully",
		}))
	}
}

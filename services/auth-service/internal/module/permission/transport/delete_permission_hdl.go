package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/module/permission/business"
	pstr "auth-service/internal/module/permission/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeletePermissionHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid permission ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		pstr := pstr.NewMySQLStorage(db)
		biz := business.NewDeletePermissionBiz(pstr)

		err = biz.DeletePermission(c, id)
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
			"message": "Permission deleted successfully",
		}))
	}
}

func SoftDeletePermissionHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid permission ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		pstr := pstr.NewMySQLStorage(db)
		biz := business.NewDeletePermissionBiz(pstr)

		err = biz.SoftDeletePermission(c, id)
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
			"message": "Permission soft deleted successfully",
		}))
	}
}

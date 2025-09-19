package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/module/permission/business"
	"auth-service/internal/module/permission/entity"
	pstr "auth-service/internal/module/permission/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdatePermissionHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid permission ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		var permissionReq entity.PermissionUpdateRequest

		if err := c.ShouldBindJSON(&permissionReq); err != nil {
			validationErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid request format")
			c.JSON(validationErr.StatusCode(), validationErr)
			return
		}

		pstr := pstr.NewMySQLStorage(db)
		biz := business.NewUpdatePermissionBiz(pstr)

		permission, err := biz.UpdatePermission(c, id, &permissionReq)
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
			"message":    "Permission updated successfully",
			"permission": permission,
		}))
	}
}

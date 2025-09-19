package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/module/permission/business"
	"auth-service/internal/module/permission/entity"
	pstr "auth-service/internal/module/permission/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePermissionHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var permissionReq entity.PermissionRequest

		if err := c.ShouldBindJSON(&permissionReq); err != nil {
			validationErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid request format")
			c.JSON(validationErr.StatusCode(), validationErr)
			return
		}

		pstr := pstr.NewMySQLStorage(db)
		biz := business.NewCreatePermissionBiz(pstr)

		permission, err := biz.CreateNewPermission(c, &permissionReq)
		if err != nil {
			if defaultErr, ok := err.(*common.DefaultError); ok {
				c.JSON(defaultErr.StatusCode(), defaultErr)
			} else {
				internalErr := common.ErrInternalServerError.WithTrace(err)
				c.JSON(internalErr.StatusCode(), internalErr)
			}
			return
		}

		c.JSON(http.StatusCreated, common.ResponseData(map[string]interface{}{
			"message":       "Permission created successfully",
			"permission_id": permission.Id,
		}))
	}
}

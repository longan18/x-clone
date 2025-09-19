package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/business"
	"auth-service/internal/module/role/entity"
	rstr "auth-service/internal/module/role/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRoleHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var roleReq entity.RoleRequest

		if err := c.ShouldBindJSON(&roleReq); err != nil {
			validationErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid request format")
			c.JSON(validationErr.StatusCode(), validationErr)
			return
		}

		rstr := rstr.NewMySQLStorage(db)
		biz := business.NewCreateRoleBiz(rstr)

		role, err := biz.CreateNewRole(c, &roleReq)
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
			"message": "Role created successfully",
			"role_id": role.Id,
		}))
	}
}

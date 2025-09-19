package transport

import (
	"auth-service/internal/common"
	"auth-service/internal/module/role/business"
	"auth-service/internal/module/role/entity"
	rstr "auth-service/internal/module/role/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateRoleHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid role ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		var roleReq entity.RoleUpdateRequest

		if err := c.ShouldBindJSON(&roleReq); err != nil {
			validationErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid request format")
			c.JSON(validationErr.StatusCode(), validationErr)
			return
		}

		rstr := rstr.NewMySQLStorage(db)
		biz := business.NewUpdateRoleBiz(rstr)

		role, err := biz.UpdateRole(c, id, &roleReq)
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
			"message": "Role updated successfully",
			"role":    role,
		}))
	}
}

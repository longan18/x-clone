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

func ListRoleHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitStr := c.DefaultQuery("limit", "10")
		offsetStr := c.DefaultQuery("offset", "0")

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit <= 0 {
			limit = 10
		}

		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			offset = 0
		}

		rstr := rstr.NewMySQLStorage(db)
		biz := business.NewListRoleBiz(rstr)

		roles, count, err := biz.ListRoles(limit, offset)
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
			"roles": roles,
			"pagination": map[string]interface{}{
				"total":  count,
				"limit":  limit,
				"offset": offset,
			},
		}))
	}
}

func GetRoleHdl(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		strId := c.Param("id")

		id, err := strconv.Atoi(strId)
		if err != nil {
			badRequestErr := common.ErrBadRequest.WithTrace(err).WithReason("Invalid role ID format")
			c.JSON(badRequestErr.StatusCode(), badRequestErr)
			return
		}

		rstr := rstr.NewMySQLStorage(db)
		biz := business.NewGetRoleBiz(rstr)

		role, err := biz.GetRoleById(id)
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
			"role": role,
		}))
	}
}

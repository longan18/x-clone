package initiative

import (
	"auth-service/global"
	"auth-service/internal/common"
	usertransp "auth-service/internal/module/user/transport"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.GET("", usertransp.ListUserTransp(db))
			user.POST("/store", usertransp.CreateUserTransp(db))
			user.GET("/:id", usertransp.GetUserTransp(db))
			user.PUT("/:id/update", usertransp.UpdateUserTransp(db))
			user.DELETE("/:id/delete", usertransp.DeleteUserTransp(db))
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, common.ResponseData(map[string]interface{}{
			"message": "Ping to auth service successful.",
		}))
	})

	port := fmt.Sprintf(":%d", global.Config.Server.Port)

	r.Run(port)
}

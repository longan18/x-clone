package initiative

import (
	"auth-service/global"
	"auth-service/internal/common"
	userTransport "auth-service/internal/module/user/transport"
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
			user.GET("", userTransport.ListUserHdl(db))
			user.POST("/store", userTransport.CreateUserHdl(db))
			user.GET("/:id", userTransport.GetUserHdl(db))
			user.PUT("/:id/update", userTransport.UpdateUserHdl(db))
			user.DELETE("/:id/delete", userTransport.DeleteUserHdl(db))
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

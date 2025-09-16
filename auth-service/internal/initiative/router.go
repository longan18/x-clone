package initiative

import (
	"auth-service/global"
	usertransp "auth-service/internal/transport/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		user := v1.Group("/user")
		{
			user.GET("", usertransp.ListUserTransp)
			user.POST("/store", usertransp.CreateUserTransp)
			user.GET("/:id", usertransp.GetUserTransp)
			user.PUT("/:id/update", usertransp.UpdateUserTransp)
			user.DELETE("/:id/delete", usertransp.DeleteUserTransp)
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ping to auth service successful.",
		})
	})

	port := fmt.Sprintf(":%d", global.Config.Server.Port)

	r.Run(port)
}

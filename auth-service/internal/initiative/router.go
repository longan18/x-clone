package initiative

import (
	"fmt"
	"net/http"
	"auth-service/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ping to auth service successful.",
		})
	})

	port := fmt.Sprintf(":%d", global.Config.Server.Port)

	r.Run(port)
}

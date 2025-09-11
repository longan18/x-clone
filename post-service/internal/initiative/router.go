package initiative

import (
	"fmt"
	"net/http"
	"post-service/global"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ping to post service successful",
		})
	})

	port := fmt.Sprintf(":%d", global.Config.Server.Port)

	r.Run(port)
}

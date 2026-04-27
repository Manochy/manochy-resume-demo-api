package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//InitMongo()

	r := gin.Default()

	r.POST("/login", Login)

	// health check
	r.GET("/check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"status": true,
		})
	})

	auth := r.Group("/")
	auth.Use(JWTAuth())

	auth.GET("/members", GetMembers)
	auth.POST("/members", CreateMember)
	auth.PUT("/members/:id", UpdateMember)
	auth.DELETE("/members/:id", DeleteMember)

	auth.POST("/chatbot", Chatbot)

	return r
}

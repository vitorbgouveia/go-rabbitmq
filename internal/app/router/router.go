package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorbgouveia/go-rabbitmq/internal/presentation/controllers"
)

// ConfigRoutes load http routess
func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("app/v1")
	{
		books := main.Group("rabbitmq")
		{
			books.POST("/hello-word", controllers.HelloWord)
			books.POST("/pub-sub", controllers.HelloWord)
		}
	}

	return router
}

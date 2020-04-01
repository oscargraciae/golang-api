package router

import (
	"api-app/api/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter config api routes
func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", controllers.GetUsers)
			users.POST("/", controllers.CreateUser)
			users.POST("/login", controllers.Login)
		}
	}
}

package routes

import (
	"web-server/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", controllers.GetHello)

	grpUser := r.Group("/user")
	{
		grpUser.GET("/", controllers.GetUsers)
		grpUser.POST("/", controllers.CreateUser)
		grpUser.GET("/:id", controllers.GetUserByID)
		grpUser.PUT("/:id", controllers.UpdateUser)
		grpUser.DELETE("/:id", controllers.DeleteUser)
	}
	return r
}

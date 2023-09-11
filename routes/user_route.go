package routes

import (
	"github.com/Badiru27/hngx-stage-2/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(eng *gin.Engine) {
	eng.POST("/api", controllers.CreateUser)
	eng.GET("/api/:userId", controllers.GetUser)
	eng.PUT("/api/:userId", controllers.UpdateUser)
	eng.DELETE("/api/:userId", controllers.DeleteUser)
}

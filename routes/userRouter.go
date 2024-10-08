package routes

import (
	controller "restuarant_app/controllers"

	"github.com/gin-gonic/gin"
)


func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
	incomingRoutes.POST("/users/singup", controller.SingUp)
	incomingRoutes.POST("/users/login", controller.Login())
}
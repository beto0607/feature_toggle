package routes

import (
	"github.com/gin-gonic/gin"
	"toggler/controllers"
)

func UserRoute(router *gin.RouterGroup) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
}

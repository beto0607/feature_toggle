package routes

import (
	"github.com/gin-gonic/gin"
	"toggler/controllers"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetAUser())
}

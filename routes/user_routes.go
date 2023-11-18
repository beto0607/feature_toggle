package routes

import (
	"github.com/gin-gonic/gin"
	"toggler/controllers"
)

func UsersRoute(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/users", controllers.CreateUser())
	routerGroup.GET("/users/:userId", controllers.GetAUser())
}

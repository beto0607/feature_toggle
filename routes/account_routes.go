package routes

import (
	"toggler/controllers"

	"github.com/gin-gonic/gin"
)

func AccountsRoutes(routeGroup *gin.RouterGroup) {
	routeGroup.GET("/accounts/", controllers.GetAccounts())
	routeGroup.POST("/accounts/", controllers.CreateAccount())
	routeGroup.GET("/accounts/:accountId", controllers.GetAccount())
	routeGroup.PUT("/accounts/:accountId", controllers.EditAccount())
	routeGroup.PATCH("/accounts/:accountId", controllers.EditAccount())
	routeGroup.DELETE("/accounts/:accountId", controllers.DeleteAccount())
}

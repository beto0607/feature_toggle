package routes

import (
	"github.com/gin-gonic/gin"
)

func DoApiRouting(routerGroup *gin.RouterGroup) {
	UsersRoute(routerGroup)
	FeaturesRoutes(routerGroup)
	AccountsRoutes(routerGroup)
}

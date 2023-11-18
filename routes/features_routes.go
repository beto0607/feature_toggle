package routes

import (
	"toggler/controllers"

	"github.com/gin-gonic/gin"
)

func FeaturesRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/features/:accountId/", controllers.GetFeatures())
	routerGroup.POST("/features/:accountId/", controllers.CreateFeature())
	routerGroup.GET("/features/:accountId/:featureId", controllers.GetFeature())
	routerGroup.PATCH("/features/:accountId/:featureId", controllers.EditFeature())
	routerGroup.PUT("/features/:accountId/:featureId", controllers.EditFeature())
	routerGroup.DELETE("/features/:accountId/:featureId", controllers.DeleteFeature())
}

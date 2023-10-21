package routes

import (
	"toggler/controllers"

	"github.com/gin-gonic/gin"
)

func FeatureRoutes(router *gin.RouterGroup) {
	router.GET("/feature/:accountId/", controllers.GetFeatures())
	router.POST("/feature/:accountId/", controllers.CreateFeature())
	router.GET("/feature/:accountId/:featureId", controllers.GetFeature())
	router.PATCH("/feature/:accountId/:featureId", controllers.EditFeature())
	router.PUT("/feature/:accountId/:featureId", controllers.EditFeature())
	router.DELETE("/feature/:accountId/:featureId", controllers.DeleteFeature())
}

package controllers

import (
	"log"
	"net/http"
	"toggler/data"
	togglerError "toggler/error"
	"toggler/models"
	"toggler/responses"

	"github.com/gin-gonic/gin"
)

func GetFeatures() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId := c.Param("accountId")
		var features []models.Feature

		account, accountErr := data.GetAccount(accountId)
		if accountErr != nil {
			togglerError.SendException(c, accountErr.Error(), "Account not found")
			return
		}

		features, featuresErr := data.GetFeatures(account.Id)
		if featuresErr != nil {
			togglerError.SendException(c, featuresErr.Error(), "Couldn't load features")
			return
		}

		c.JSON(
			http.StatusOK,
			responses.FeaturesResponse{
				Data: features,
			},
		)
	}
}
func GetFeature() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateFeature() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountId := c.Param("accountId")

		account, accountErr := data.GetAccount(accountId)
		if accountErr != nil {
			togglerError.SendException(c, accountErr.Error(), "Account not found")
			return
		}
		var feature models.Feature

		if err := c.BindJSON(&feature); err != nil {
			log.Println(err.Error())
			togglerError.SendException(c, togglerError.BadRequest, "")
			return
		}
		if validationErr := validate.Struct(&feature); validationErr != nil {
			log.Println(validationErr.Error())
			togglerError.SendException(c, togglerError.BadRequest, "Invalid data")
			return
		}

		feature.AccountId = account.Id.String()

		newFeature, newFeatureErr := data.AddFeature(&feature)

		if newFeatureErr != nil {
			log.Println(newFeatureErr.Error())
			togglerError.SendException(c, togglerError.InternalError, "Couldn't save feature")
			return
		}

		c.JSON(
			http.StatusCreated,
			responses.FeatureResponse{
				Data: *newFeature,
			},
		)

	}
}

func EditFeature() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func DeleteFeature() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

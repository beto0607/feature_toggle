package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	"toggler/models"
)

func GetFeatures(w http.ResponseWriter, request *http.Request) {

	// features, featuresErr := data.GetFeatures()
	// if featuresErr != nil {
	// 	w.WriteHeader(http.InternalError)
	// 	if status == http.StatusNotFound {
	// 		fmt.Fprint(w, "custom 404")
	// 	}
	//
	// 	togglerError.SendException(c, featuresErr.Error(), "Couldn't load features")
	// 	return
	// }

	newFeature := models.Feature{
		AccountId: "account",
		Name:      "name",
		Enabled:   true,
		CreatedBy: time.Now().UTC().String(),
		CreatedAt: time.Now().UTC().String(),
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&newFeature)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// request.JSON(
	// 	http.StatusOK,
	// 	responses.FeaturesResponse{
	// 		Data: features,
	// 	},
	// )
}
func GetFeature(writer http.ResponseWriter, request *http.Request) {
	// type Book struct {
	// 	Title  string `json:"title"`
	// 	Author string `json:"author"`
	// }
	//
	// book := Book{"Building Web Apps with Go", "Jeremy Saenz"}
	//
	// js, err := json.Marshal(book)
	// if err != nil {
	// 	http.Error(writer, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//
	// writer.Header().Set("Content-Type", "application/json")
	// writer.Write(js)

}

// func CreateFeature() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		accountId := c.Param("accountId")
//
// 		account, accountErr := data.GetAccount(accountId)
// 		if accountErr != nil {
// 			togglerError.SendException(c, accountErr.Error(), "Account not found")
// 			return
// 		}
// 		var feature models.Feature
//
// 		if err := c.BindJSON(&feature); err != nil {
// 			log.Println(err.Error())
// 			togglerError.SendException(c, togglerError.BadRequest, "")
// 			return
// 		}
// 		if validationErr := validate.Struct(&feature); validationErr != nil {
// 			log.Println(validationErr.Error())
// 			togglerError.SendException(c, togglerError.BadRequest, "Invalid data")
// 			return
// 		}
//
// 		feature.AccountId = account.Id.String()
//
// 		newFeature, newFeatureErr := data.AddFeature(&feature)
//
// 		if newFeatureErr != nil {
// 			log.Println(newFeatureErr.Error())
// 			togglerError.SendException(c, togglerError.InternalError, "Couldn't save feature")
// 			return
// 		}
//
// 		c.JSON(
// 			http.StatusCreated,
// 			responses.FeatureResponse{
// 				Data: *newFeature,
// 			},
// 		)
//
// 	}
// }
//
// func EditFeature() gin.HandlerFunc {
// 	return func(c *gin.Context) {}
// }
//
// func DeleteFeature() gin.HandlerFunc {
// 	return func(c *gin.Context) {}
// }

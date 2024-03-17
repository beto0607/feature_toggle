package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"toggler/data"
	"toggler/models"
)

func GetFeatures(w http.ResponseWriter, request *http.Request) {
	features, featuresErr := data.GetFeatures()
	if !featuresErr {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&features)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func GetFeature(w http.ResponseWriter, r *http.Request) {
	featureId := r.PathValue("id")
	if len(featureId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	feature, ok := data.GetFeature(featureId)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&feature)
}

func CreateFeature(w http.ResponseWriter, r *http.Request) {
	var feature models.Feature

	err := json.NewDecoder(r.Body).Decode(&feature)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if validationErr := validate.Struct(&feature); validationErr != nil {
		log.Println(validationErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newFeature, newFeatureOk := data.AddFeature(&feature)

	if !newFeatureOk {
		log.Println("Couln't insert document")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newFeature)
}

func EditFeature(w http.ResponseWriter, r *http.Request) {
	featureId := r.PathValue("id")
	if len(featureId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var featureDto models.FeatureDto
	err := json.NewDecoder(r.Body).Decode(&featureDto)
	log.Println(featureDto.Flags)
	log.Println(featureDto.Enabled)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if validationErr := validate.Struct(&featureDto); validationErr != nil {
		log.Println(validationErr.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedFeature, ok := data.EditFeature(featureId, featureDto)
	if !ok {
		log.Println("Couln't update document")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&updatedFeature)
}

func DeleteFeature(w http.ResponseWriter, r *http.Request) {
	featureId := r.PathValue("id")
	if len(featureId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ok := data.DeleteFeature(featureId)
	if !ok {
		log.Println("Couln't delete document")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

// func DeleteFeature() gin.HandlerFunc {
// 	return func(c *gin.Context) {}
// }

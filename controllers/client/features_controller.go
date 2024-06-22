package client

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"toggler/data"
	"toggler/models"
	"toggler/utils"
)

func FeaturesList(w http.ResponseWriter, r *http.Request) {
	features, featuresErr := data.GetFeatures()
	if !featuresErr {
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	templ := template.Must(template.ParseFiles("templates/index.html"))

	data := map[string][]models.Feature{
		"Features": features,
	}

	templ.ExecuteTemplate(w, "features-list", data)
}

func CreateFeature(w http.ResponseWriter, r *http.Request) {
	var feature models.Feature
	err := json.NewDecoder(r.Body).Decode(&feature)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if validationErr := utils.Validator.Struct(&feature); validationErr != nil {
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
	templ := template.Must(template.ParseFiles("templates/index.html"))

	templ.ExecuteTemplate(w, "features-list-item", newFeature)
}

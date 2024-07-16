package client

import (
	"html/template"
	"log"
	"net/http"
	"time"
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
	time.Sleep(5 * time.Second)
	featureName := (r.FormValue("name"))
	featureEnabled := (r.FormValue("enabled")) == "on"
	feature := models.Feature{
		Name:    featureName,
		Enabled: featureEnabled,
		Flags:   map[string]interface{}{},
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
	log.Println(("Feature inserted"))
	templ := template.Must(template.ParseFiles("templates/index.html"))

	templ.ExecuteTemplate(w, "features-list-item", newFeature)
}

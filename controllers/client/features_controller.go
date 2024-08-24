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

	template, templateErr := utils.NewTemplate("list")

	if templateErr != nil {
		log.Panic(templateErr)
		w.WriteHeader(http.StatusInternalServerError)
	}

	data := map[string][]models.Feature{
		"Features": features,
	}

	err := template.ExecuteTemplate(w, "features-list", data)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
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

	err := templ.ExecuteTemplate(w, "features-list-item", newFeature)
	if err != nil {
		log.Print(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func UpdateFlag(w http.ResponseWriter, r *http.Request) {
	formErr := r.ParseForm()
	featureId := r.PathValue("id")
	flagName := r.PathValue("flagName")
	if len(featureId) == 0 || len(flagName) == 0 || formErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	feature, ok := data.GetFeature(featureId)
	if !ok || feature.Flags[flagName] == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	newValue := r.FormValue(flagName)
	feature.Flags[flagName] = newValue
	featureDto := models.FeatureDto{
		Flags: &feature.Flags,
	}
	_, ok = data.EditFeature(featureId, featureDto)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ := template.Must(template.ParseFiles("templates/features/flag-string-list-item.html"))
	params := [3]any{featureId, flagName, newValue}
	err := templ.ExecuteTemplate(w, "flag-string-list-item", params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ToggleFlag(w http.ResponseWriter, r *http.Request) {
	featureId := r.PathValue("id")
	flagName := r.PathValue("flagName")
	if len(featureId) == 0 || len(flagName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	feature, ok := data.GetFeature(featureId)
	if !ok || feature.Flags[flagName] == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	newValue := feature.Flags[flagName] == false
	feature.Flags[flagName] = newValue
	featureDto := models.FeatureDto{
		Flags: &feature.Flags,
	}

	_, ok = data.EditFeature(featureId, featureDto)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	templ := template.Must(template.ParseFiles("templates/features/flag-boolean-list-item.html"))
	params := [3]any{featureId, flagName, newValue}
	err := templ.ExecuteTemplate(w, "flag-boolean-list-item", params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func ToggleFeature(w http.ResponseWriter, r *http.Request) {
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

	newValue := !feature.Enabled
	feature.Enabled = newValue
	featureDto := models.FeatureDto{
		Enabled: &newValue,
	}

	_, ok = data.EditFeature(featureId, featureDto)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	template, templateError := utils.NewTemplate("features-list-item-status")
	if templateError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := template.ExecuteTemplate(w, "features-list-item-status", feature)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

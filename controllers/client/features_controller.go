package client

import (
	"html/template"
	"net/http"
	"toggler/data"
	"toggler/models"
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

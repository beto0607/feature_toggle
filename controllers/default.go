package controllers

import (
	"log"
	"net/http"
	"toggler/data"
	"toggler/models"
	"toggler/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	features, featuresErr := data.GetFeatures()
	if !featuresErr {
		log.Panic(featuresErr)
	}

	template, templateErr := utils.NewTemplate("index")

	if templateErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Panic(templateErr)
	}

	data := map[string][]models.Feature{
		"Features": features,
	}
	err := template.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Panic(err)
	}
}

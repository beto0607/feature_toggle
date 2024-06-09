package routes

import (
	"log"
	"net/http"
	"toggler/controllers"
)

type Route struct {
	handler http.HandlerFunc
	method  string
	path    string
}

func DoApiRouting() *http.ServeMux {
    log.Println("Doing routing...")
	featuresRouter := featureAPIRouting()

	apiRouter := http.NewServeMux()
	apiRouter.Handle("/api/", http.StripPrefix("api", featuresRouter))

	return apiRouter
}

func featureAPIRouting() *http.ServeMux {
	featuresRouter := http.NewServeMux()
	featuresRouter.HandleFunc("GET /features", controllers.GetFeatures)
	featuresRouter.HandleFunc("GET /features/{id}", controllers.GetFeature)
	featuresRouter.HandleFunc("POST /features", controllers.CreateFeature)
	featuresRouter.HandleFunc("PUT /features/{id}", controllers.EditFeature)
	featuresRouter.HandleFunc("PATCH /features/{id}", controllers.EditFeature)
	featuresRouter.HandleFunc("DELETE /features/{id}", controllers.DeleteFeature)
    log.Println("Features API added")
	return featuresRouter
}

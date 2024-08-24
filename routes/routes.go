package routes

import (
	"log"
	"net/http"
	"toggler/controllers"
	"toggler/controllers/api"
	"toggler/controllers/client"
)

func DoRouting() *http.ServeMux {
	defaultRouter := doDefaultRouting()
	apiRouter := doApiRouting()
	clientRouter := doClientRouting()
	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter))
	router.Handle("/client/", http.StripPrefix("/client", clientRouter))
	router.Handle("/", defaultRouter)

	return router
}

func doDefaultRouting() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", controllers.Index)
	return router
}

func doApiRouting() *http.ServeMux {
	log.Println("Doing API routing...")
	router := http.NewServeMux()
	featureAPIRouting(router)

	return router
}

func featureAPIRouting(router *http.ServeMux) {
	router.HandleFunc("GET /features", api.GetFeatures)
	router.HandleFunc("GET /features/{id}", api.GetFeature)
	router.HandleFunc("POST /features", api.CreateFeature)
	router.HandleFunc("PUT /features/{id}", api.EditFeature)
	router.HandleFunc("PATCH /features/{id}", api.EditFeature)
	router.HandleFunc("DELETE /features/{id}", api.DeleteFeature)
	log.Println("Features API added")
}

func doClientRouting() *http.ServeMux {
	log.Println("Doing Client routing...")
	router := http.NewServeMux()
	featuresClientRouting(router)

	return router

}
func featuresClientRouting(router *http.ServeMux) {
	router.HandleFunc("GET /features", client.FeaturesList)
	router.HandleFunc("POST /features", client.CreateFeature)
	router.HandleFunc("DELETE /features/{id}", client.DeleteFeature)
	router.HandleFunc("PUT /features/{id}/toggle", client.ToggleFeature)
	router.HandleFunc("PUT /features/{id}/flags/{flagName}/value", client.UpdateFlag)
	router.HandleFunc("PUT /features/{id}/flags/{flagName}/toggle", client.ToggleFlag)
	log.Println("Features Client added")
}

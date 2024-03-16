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

func DoApiRouting() {
	featureRouting()
}

func featureRouting() {
	var routes = []Route{
		{
			method:  "GET",
			path:    "/api/features/",
			handler: controllers.GetFeatures,
		},
		{
			method:  "GET",
			path:    "/api/features/{id}",
			handler: controllers.GetFeature,
		},
	}

	for i := 0; i < len(routes); i++ {
		route := routes[i]

		var pattern = route.path
		if len(route.method) > 0 {
			pattern = route.method + " " + route.path
		}
		http.HandleFunc(pattern, route.handler)
		log.Println(pattern + " was added")
	}
}

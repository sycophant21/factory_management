package main

import (
	controllers "factory_management_go/app/controller/data"
	"factory_management_go/app/repository"
	"factory_management_go/app/util"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()
	var locationController = controllers.LocationController{}
	locationController.Initialise()
	mux.Handle("/location/", contextPathMiddleware("/location", locationController.LocationMutex))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func contextPathMiddleware(contextPath string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request path starts with the desired context path
		if !strings.HasPrefix(r.URL.Path, contextPath) {
			http.NotFound(w, r)
			return
		}

		// Remove the context path prefix from the URL path
		r.URL.Path = strings.TrimPrefix(r.URL.Path, contextPath)
		next.ServeHTTP(w, r) // Pass to the next handler
	})
}

func init() {
	err := util.LoadProperties()
	if err != nil {
		log.Fatal(err)
	}
	err = repository.InitialiseEngine()
	if err != nil {
		log.Fatal()
	}
}

// need to change everything from functional to object based (object of controller, service, repository etc)

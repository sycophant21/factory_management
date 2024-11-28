package main

import (
	controllers "factory_management_go/app/controller/data"
	"factory_management_go/app/domain/dao/location"
	engine "factory_management_go/app/engine"
	logg "factory_management_go/app/log"
	repository "factory_management_go/app/repository"
	services "factory_management_go/app/service/data"
	"factory_management_go/app/util"
	"log"
	"net/http"
	"strings"
)

func main() {
	locationRepoEngine, err := engine.InitialiseEngine[location.Location]()
	if err != nil {
		logg.Logger.Error(err.Error())
	}
	locationTypeRepoEngine, err := engine.InitialiseEngine[location.LocationType]()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	locationTypeController, locationTypeService, locationTypeRepository, err := initLocationType(locationTypeRepoEngine)
	if err != nil {
		log.Fatal(err)
	}
	locationController, _, _, err := initLocation(locationRepoEngine, locationTypeRepository, locationTypeService)
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/locationType/", contextPathMiddleware("/locationType", locationTypeController.Mutex))
	mux.Handle("/location/", contextPathMiddleware("/location", locationController.Mutex))
	logg.Logger.Error(http.ListenAndServe(":8080", mux).Error())
}

func contextPathMiddleware(contextPath string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, contextPath) {
			http.NotFound(w, r)
			return
		}
		r.URL.Path = strings.TrimPrefix(r.URL.Path, contextPath)
		next.ServeHTTP(w, r) // Pass to the next handler
	})
}

func init() {
	err := util.LoadProperties()
	if err != nil {
		log.Fatal(err)
	}
	err = logg.Initialise("logs.txt")
	if err != nil {
		log.Fatal(err)
	}
}
func initLocationType(eng *engine.RepoEngine[location.LocationType]) (*controllers.LocationTypeController, *services.LocationTypeService, *repository.LocationTypeRepository, error) {
	var locationTypeRepository = repository.LocationTypeRepository{Eng: eng}
	var locationTypeService = services.LocationTypeService{Repository: &locationTypeRepository}
	var locationTypeController = controllers.LocationTypeController{Service: &locationTypeService}
	locationTypeController.Initialise()
	return &locationTypeController, &locationTypeService, &locationTypeRepository, nil
}
func initLocation(eng *engine.RepoEngine[location.Location], locationTypeRepository *repository.LocationTypeRepository, locationTypeService *services.LocationTypeService) (*controllers.LocationController, *services.LocationService, *repository.LocationRepository, error) {
	var locationRepository = repository.LocationRepository{Eng: eng, Ltr: locationTypeRepository}
	var locationService = services.LocationService{Repository: &locationRepository, LocationTypeService: locationTypeService}
	var locationController = controllers.LocationController{Service: &locationService}
	locationController.Initialise()
	return &locationController, &locationService, &locationRepository, nil
}

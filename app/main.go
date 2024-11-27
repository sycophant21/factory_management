package main

import (
	controllers "factory_management_go/app/controller/data"
	"factory_management_go/app/domain/dao/location"
	logg "factory_management_go/app/log"
	repo "factory_management_go/app/repository"
	repos "factory_management_go/app/repository/data"
	services "factory_management_go/app/service/data"
	"factory_management_go/app/util"
	"log"
	"net/http"
	"strings"
)

func main() {
	locationRepoEngine, err := repo.InitialiseEngine[location.Location]()
	if err != nil {
		logg.Logger.Error(err.Error(), "app.main")
		//log.Fatal(err)
	}
	locationTypeRepoEngine, err := repo.InitialiseEngine[location.LocationType]()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	locationTypeController, locationTypeService, _, err := initLocationType(locationTypeRepoEngine)
	if err != nil {
		log.Fatal(err)
	}
	locationController, _, _, err := initLocation(locationRepoEngine, locationTypeService)
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/locationType/", contextPathMiddleware("/locationType", locationTypeController.Mutex))
	mux.Handle("/location/", contextPathMiddleware("/location", locationController.Mutex))
	logg.Logger.Error(http.ListenAndServe(":8080", mux).Error(), "app.main")
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
func initLocationType(eng *repo.RepoEngine[location.LocationType]) (*controllers.LocationTypeController, *services.LocationTypeService, *repos.LocationTypeRepository, error) {
	var locationTypeRepository = repos.LocationTypeRepository{Eng: eng}
	var locationTypeService = services.LocationTypeService{Repository: &locationTypeRepository}
	var locationTypeController = controllers.LocationTypeController{Service: &locationTypeService}
	locationTypeController.Initialise()
	return &locationTypeController, &locationTypeService, &locationTypeRepository, nil
}
func initLocation(eng *repo.RepoEngine[location.Location], locationTypeService *services.LocationTypeService) (*controllers.LocationController, *services.LocationService, *repos.LocationRepository, error) {
	var locationRepository = repos.LocationRepository{Eng: eng}
	var locationService = services.LocationService{Repository: &locationRepository, LocationTypeService: locationTypeService}
	var locationController = controllers.LocationController{Service: &locationService}
	locationController.Initialise()
	return &locationController, &locationService, &locationRepository, nil
}

package main

import (
	initialiser "factory_management_go/app/init"
	logg "factory_management_go/app/log"
	"factory_management_go/app/middleware"
	"factory_management_go/app/util"
	"log"
	"net/http"
	"os"
)

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

func main() {
	locationTypeController, locationTypeService, locationTypeRepository, err := initialiser.InitLocationType()
	if err != nil {
		logg.Logger.Error(err.Error())
		os.Exit(1)
	}
	locationController, locationService, _, err := initialiser.InitLocation(locationTypeRepository)
	if err != nil {
		logg.Logger.Error(err.Error())
		os.Exit(1)
	}
	componentTypeController, componentTypeService, componentTypeRepository, err := initialiser.InitComponentType()
	if err != nil {
		logg.Logger.Error(err.Error())
		os.Exit(1)
	}
	componentController, _, _, err := initialiser.InitComponent(componentTypeRepository)
	if err != nil {
		logg.Logger.Error(err.Error())
		os.Exit(1)
	}
	optionsController, _, err := initialiser.InitOptions(locationTypeService, locationService, componentTypeService)
	if err != nil {
		logg.Logger.Error(err.Error())
		os.Exit(1)
	}
	mux := http.NewServeMux()
	mux.Handle("/locationType/", middleware.ContextPathMiddleware("/locationType", locationTypeController.Mutex))
	mux.Handle("/location/", middleware.ContextPathMiddleware("/location", locationController.Mutex))
	mux.Handle("/component/", middleware.ContextPathMiddleware("/component", componentController.Mutex))
	mux.Handle("/spareType/", middleware.ContextPathMiddleware("/spareType", componentController.Mutex))
	mux.Handle("/componentType/", middleware.ContextPathMiddleware("/componentType", componentTypeController.Mutex))
	mux.Handle("/options/", middleware.ContextPathMiddleware("/options", optionsController.Mutex))
	logg.Logger.Error(http.ListenAndServe(":8080", mux).Error())
}

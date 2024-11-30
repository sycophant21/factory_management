package init

import (
	dataControllers "factory_management_go/app/controller/data"
	optionsController "factory_management_go/app/controller/option"
	"factory_management_go/app/domain/dao/component"
	location "factory_management_go/app/domain/dao/location"
	engine "factory_management_go/app/engine"
	logg "factory_management_go/app/log"
	repository "factory_management_go/app/repository"
	dataServices "factory_management_go/app/service/data"
	optionsService "factory_management_go/app/service/option"
)

func InitLocationType() (*dataControllers.LocationTypeController, *dataServices.LocationTypeService, *repository.LocationTypeRepository, error) {
	eng, err := engine.InitialiseEngine[location.LocationType]()
	if err != nil {
		logg.Logger.Error(err.Error())
		return nil, nil, nil, err
	}
	var locationTypeRepository = repository.LocationTypeRepository{Eng: eng}
	var locationTypeService = dataServices.LocationTypeService{LocationTypeRepository: &locationTypeRepository}
	var locationTypeController = dataControllers.LocationTypeController{LocationTypeService: &locationTypeService}
	locationTypeController.Initialise()
	return &locationTypeController, &locationTypeService, &locationTypeRepository, nil
}
func InitLocation(locationTypeRepository *repository.LocationTypeRepository) (*dataControllers.LocationController, *dataServices.LocationService, *repository.LocationRepository, error) {
	eng, err := engine.InitialiseEngine[location.Location]()
	if err != nil {
		logg.Logger.Error(err.Error())
		return nil, nil, nil, err
	}
	var locationRepository = repository.LocationRepository{Eng: eng, LocationTypeRepository: locationTypeRepository}
	var locationService = dataServices.LocationService{LocationRepository: &locationRepository}
	var locationController = dataControllers.LocationController{LocationService: &locationService}
	locationController.Initialise()
	return &locationController, &locationService, &locationRepository, nil
}

func InitComponentType() (*dataControllers.ComponentTypeController, *dataServices.ComponentTypeService, *repository.ComponentTypeRepository, error) {
	eng, err := engine.InitialiseEngine[component.ComponentType]()
	if err != nil {
		logg.Logger.Error(err.Error())
		return nil, nil, nil, err
	}
	var componentTypeRepository = repository.ComponentTypeRepository{Eng: eng}
	var componentTypeService = dataServices.ComponentTypeService{ComponentTypeRepository: &componentTypeRepository}
	var componentTypeController = dataControllers.ComponentTypeController{ComponentTypeService: &componentTypeService}
	componentTypeController.Initialise()
	return &componentTypeController, &componentTypeService, &componentTypeRepository, nil
}

func InitComponent(componentTypeRepository *repository.ComponentTypeRepository) (*dataControllers.ComponentController, *dataServices.ComponentService, *repository.ComponentRepository, error) {
	eng, err := engine.InitialiseEngine[component.Component]()
	if err != nil {
		logg.Logger.Error(err.Error())
		return nil, nil, nil, err
	}
	var componentRepository = repository.ComponentRepository{Eng: eng, ComponentTypeRepository: componentTypeRepository}
	var componentService = dataServices.ComponentService{ComponentRepository: &componentRepository}
	var componentController = dataControllers.ComponentController{ComponentService: &componentService}
	componentController.Initialise()
	return &componentController, &componentService, &componentRepository, nil
}

func InitOptions(locationTypeService *dataServices.LocationTypeService, locationService *dataServices.LocationService, componentTypeService *dataServices.ComponentTypeService) (*optionsController.OptionsController, *optionsService.OptionsService, error) {
	var optionService = optionsService.OptionsService{LocationTypeService: locationTypeService, LocationService: locationService, ComponentTypeService: componentTypeService}
	var optionController = optionsController.OptionsController{Service: &optionService}
	optionController.Initialise()
	return &optionController, &optionService, nil
}

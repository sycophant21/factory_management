package option

import (
	"encoding/json"
	option "factory_management_go/app/domain/dto/response/option"
	"factory_management_go/app/domain/dto/response/wrapper"
	logg "factory_management_go/app/log"
	service "factory_management_go/app/service/option"
	"net/http"
	"time"
)

type OptionsController struct {
	Mutex   *http.ServeMux
	Service *service.OptionsService
}

func (o *OptionsController) Initialise() {
	o.Mutex = http.NewServeMux()
	o.Mutex.HandleFunc("GET /getAllLocationTypes", o.GetAllLocationTypes)
	o.Mutex.HandleFunc("GET /getAllComponentTypes", o.GetAllComponentTypes)
}

func (o *OptionsController) GetAllLocationTypes(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	locationTypeOptions, err := o.Service.GetAllLocationTypeOptions(request.Header.Get("Company-Id"))
	if err != nil {
		logg.Logger.Error(err.Error())
		now := time.Now()
		err = json.NewEncoder(writer).Encode(option.AllOptionsResponseDto{ResponseMetadata: &wrapper.ResponseMetadata{Message: "Failure", HttpCode: 500, CreatedAt: &now, LastUpdatedAt: &now}})
	} else {
		now := time.Now()
		err = json.NewEncoder(writer).Encode(option.AllOptionsResponseDto{ResponseMetadata: &wrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: &now, LastUpdatedAt: &now}, Data: locationTypeOptions})
	}
	if err != nil {
		logg.Logger.Error(err.Error())
	}
}

func (o *OptionsController) GetAllComponentTypes(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	componentTypesOptions, err := o.Service.GetAllComponentTypeOptions(request.Header.Get("Company-Id"))
	if err != nil {
		logg.Logger.Error(err.Error())
		now := time.Now()
		err = json.NewEncoder(writer).Encode(option.AllOptionsResponseDto{ResponseMetadata: &wrapper.ResponseMetadata{Message: "Failure", HttpCode: 500, CreatedAt: &now, LastUpdatedAt: &now}, Data: componentTypesOptions})
	} else {
		now := time.Now()
		err = json.NewEncoder(writer).Encode(option.AllOptionsResponseDto{ResponseMetadata: &wrapper.ResponseMetadata{Message: "Success", HttpCode: 200, CreatedAt: &now, LastUpdatedAt: &now}, Data: componentTypesOptions})
	}
	if err != nil {
		logg.Logger.Error(err.Error())
	}
}

/*

@GetMapping("/getAllComponentUseCategories")
public AllEntityOptionResponseDto getAllComponentUseCategories(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<ComponentUseCategory> componentUseCategories = componentUseCategoryService.getAllComponentUseCategories(companyId);
    List<EntityOptionResponseDto> componentUseCategoryResponseDtos = new ArrayList<>();
    int index = 1;
    for (ComponentUseCategory componentUseCategory : componentUseCategories) {
        componentUseCategoryResponseDtos.add(DaoToDto.convertComponentUseCategoryToEntityOptionResponseDto(index++, componentUseCategory));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentUseCategoryResponseDtos);
}

@GetMapping("/getAllComponentReturnTypes")
public AllEntityOptionResponseDto getAllComponentReturnTypes(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<ComponentReturnType> componentReturnTypes = componentReturnTypeService.getAllComponentReturnTypes(companyId);
    List<EntityOptionResponseDto> componentReturnTypeResponseDtos = new ArrayList<>();
    int index = 1;
    for (ComponentReturnType componentReturnType : componentReturnTypes) {
        componentReturnTypeResponseDtos.add(DaoToDto.convertComponentReturnTypeToEntityOptionResponseDto(index++, componentReturnType));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentReturnTypeResponseDtos);
}

@GetMapping("/getAllComponentPositions")
public AllEntityOptionResponseDto getAllComponentPositions(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<ComponentPosition> componentPositions = componentPositionService.getAllComponentPositions(companyId);
    List<EntityOptionResponseDto> componentPositionResponseDtos = new ArrayList<>();
    int index = 1;
    for (ComponentPosition componentPosition : componentPositions) {
        componentPositionResponseDtos.add(DaoToDto.convertComponentPositionToEntityOptionResponseDto(index++, componentPosition));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentPositionResponseDtos);
}

@GetMapping("/getAllComponentUnits")
public AllEntityOptionResponseDto getAllComponentUnits(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<ComponentUnit> componentUnits = componentUnitService.getAllComponentUnit(companyId);
    List<EntityOptionResponseDto> componentUnitResponseDtos = new ArrayList<>();
    int index = 1;
    for (ComponentUnit componentUnit : componentUnits) {
        componentUnitResponseDtos.add(DaoToDto.convertComponentUnitToEntityOptionResponseDto(index++, componentUnit));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentUnitResponseDtos);
}

@GetMapping("/getAllGSTPercentages")
public AllEntityOptionResponseDto getAllGSTPercentages(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<GSTPercentage> gstPercentages = gstPercentageService.getAllGSTPercentages(companyId);
    List<EntityOptionResponseDto> gstPercentageResponseDtos = new ArrayList<>();
    int index = 1;
    for (GSTPercentage gstPercentage : gstPercentages) {
        gstPercentageResponseDtos.add(DaoToDto.convertGSTPercentageToEntityOptionResponseDto(index++, gstPercentage));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), gstPercentageResponseDtos);
}

@GetMapping(value = {"/getAllComponents", "/getAllSpareTypes"})
public AllEntityOptionResponseDto getAllComponents(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<Component> components = componentService.getAllComponents(companyId);
    List<EntityOptionResponseDto> componentsResponseDtos = new ArrayList<>();
    int index = 1;
    for (Component component : components) {
        componentsResponseDtos.add(DaoToDto.convertComponentToEntityOptionResponseDto(index++, component));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentsResponseDtos);
}

@GetMapping(value = {"/getAllComponentInventoriesForModal", "/getAllSparePartsForModal"})
public CustomEntityOptionResponseDto getAllComponentInventoriesForModal(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<ComponentInventory> componentInventories = componentInventoryService.getAllComponentInventories(companyId);
    List<Map<String, Object>> componentsResponseDtos = new ArrayList<>();
    int index = 1;
    for (ComponentInventory componentInventory : componentInventories) {
        componentsResponseDtos.add(DaoToDto.convertComponentInventoryToCustomEntityOptionResponseDto(index++, componentInventory));
    }
    return new CustomEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentsResponseDtos);
}

@GetMapping(value = {"/getAllComponentInventoriesSoldByVendorForModal", "/getAllSparePartsSoldByVendorForModal"})
public CustomEntityOptionResponseDto getAllComponentInventoriesSoldByVendorForModal(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId, @RequestParam(value = "vendor", required = false) String vendor) {
    try {
        List<ComponentInventory> componentInventories = sparePartVendorService.getAllSparePartsFromVendor(vendor, companyId);
        List<Map<String, Object>> componentsResponseDtos = new ArrayList<>();
        int index = 1;
        for (ComponentInventory componentInventory : componentInventories) {
            componentsResponseDtos.add(DaoToDto.convertComponentInventoryToCustomEntityOptionResponseDto(index++, componentInventory));
        }
        return new CustomEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentsResponseDtos);

    } catch (InvalidInputException e) {
        return new CustomEntityOptionResponseDto("Failure", 200, LocalDateTime.now(), LocalDateTime.now());
    }
}

@GetMapping(value = {"/getAllComponentInventories", "/getAllSpareParts"})
public AllEntityOptionResponseDto getAllComponentInventories(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<ComponentInventory> componentInventories = componentInventoryService.getAllComponentInventories(companyId);
    List<EntityOptionResponseDto> componentsResponseDtos = new ArrayList<>();
    int index = 1;
    for (ComponentInventory componentInventory : componentInventories) {
        componentsResponseDtos.add(DaoToDto.convertComponentInventoryToEntityOptionResponseDto(index++, componentInventory));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentsResponseDtos);
}

@GetMapping("/getAllLocations")
public AllEntityOptionResponseDto getAllLocations(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId, @RequestParam(value = "locationType", required = false) String locationType) {
    List<Location> locations;
    if (locationType == null || locationType.isBlank()) {
        locations = locationService.getAllLocations(companyId);
    } else {
        locations = locationService.getLocationsFromLocationType(companyId, locationType);
    }
    List<EntityOptionResponseDto> locationsResponseDtos = new ArrayList<>();
    int index = 1;
    for (Location location : locations) {
        locationsResponseDtos.add(DaoToDto.convertLocationToEntityOptionResponseDto(index++, location));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), locationsResponseDtos);
}

@GetMapping("/getAllComponentLifeDurationTypes")
public AllEntityOptionResponseDto getAllComponentLifeDurationTypes(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<DurationType> durationTypes = durationTypeService.getAllDurationTypes(companyId);
    List<EntityOptionResponseDto> durationTypesResponseDtos = new ArrayList<>();
    int index = 1;
    for (DurationType durationType : durationTypes) {
        durationTypesResponseDtos.add(DaoToDto.convertDurationTypeToEntityOptionResponseDto(index++, durationType));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), durationTypesResponseDtos);
}

@GetMapping(value = {"/getAllMachineComponentInventoriesForModal", "/getAllMachineSparePartsForModal"})
public CustomEntityOptionResponseDto getAllMachineComponentInventoriesForModal(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId, @RequestParam("machineId") String machineId) {
    try {
        List<ComponentInventory> componentInventories = machineService.getMachineComponents(companyId, machineId);
        List<Map<String, Object>> componentsResponseDtos = new ArrayList<>();
        int index = 1;
        for (ComponentInventory componentInventory : componentInventories) {
            componentsResponseDtos.add(DaoToDto.convertComponentInventoryToCustomEntityOptionResponseDto(index++, componentInventory));
        }
        return new CustomEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentsResponseDtos);
    } catch (InvalidInputException e) {
        return new CustomEntityOptionResponseDto("Failure, No Machine found with machineId = " + machineId, 404, LocalDateTime.now(), LocalDateTime.now(), Collections.emptyList());
    }
}
@GetMapping(value = {"/getAllMachineComponentInventoriesFromMaintenanceRequestForModal", "/getAllMachineSparePartsFromMaintenanceRequestForModal"})
public CustomEntityOptionResponseDto getAllMachineComponentInventoriesFromMaintenanceRequestForModal(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId, @RequestParam("maintenanceRequest") String maintenanceRequestId) {
    try {
        List<ComponentInventory> componentInventories = maintenanceRequestService.getMaintenanceRequestDetails(companyId, maintenanceRequestId).getMachine().getMachineComponents();
        List<Map<String, Object>> componentsResponseDtos = new ArrayList<>();
        int index = 1;
        for (ComponentInventory componentInventory : componentInventories) {
            componentsResponseDtos.add(DaoToDto.convertComponentInventoryToCustomEntityOptionResponseDto(index++, componentInventory));
        }
        return new CustomEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), componentsResponseDtos);
    } catch (InvalidInputException e) {
        return new CustomEntityOptionResponseDto("Failure, No Maintenance Request found with maintenanceRequestId = " + maintenanceRequestId, 404, LocalDateTime.now(), LocalDateTime.now(), Collections.emptyList());
    }
}

@GetMapping("/getAllSparePartVendors")
public AllEntityOptionResponseDto getAllSparePartVendors(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<SparePartVendor> sparePartVendors = sparePartVendorService.getAllSparePartVendors(companyId);
    List<EntityOptionResponseDto> sparePartVendorsResponseDtos = new ArrayList<>();
    int index = 1;
    for (SparePartVendor sparePartVendor : sparePartVendors) {
        sparePartVendorsResponseDtos.add(DaoToDto.convertSparePartVendorToEntityOptionResponseDto(index++, sparePartVendor));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), sparePartVendorsResponseDtos);
}

@GetMapping("/getAllServiceVendors")
public AllEntityOptionResponseDto getAllServiceVendors(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<ServiceVendor> serviceVendors = serviceVendorService.getAllServiceVendors(companyId);
    List<EntityOptionResponseDto> serviceVendorsResponseDtos = new ArrayList<>();
    int index = 1;
    for (ServiceVendor serviceVendor : serviceVendors) {
        serviceVendorsResponseDtos.add(DaoToDto.convertServiceVendorToEntityOptionResponseDto(index++, serviceVendor));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), serviceVendorsResponseDtos);
}

@GetMapping("/getAllMachinesForModal")
public CustomEntityOptionResponseDto getAllMachinesForModal(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<Machine> machines = machineService.getAllMachines(companyId);
    List<Map<String, Object>> machinesResponseDtos = new ArrayList<>();
    int index = 1;
    for (Machine machine : machines) {
        machinesResponseDtos.add(DaoToDto.convertMachineToCustomEntityOptionResponseDto(index++, machine));
    }
    return new CustomEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), machinesResponseDtos);
}

@GetMapping("/getAllMachines")
public AllEntityOptionResponseDto getAllMachines(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<Machine> machines = machineService.getAllMachines(companyId);
    List<EntityOptionResponseDto> machinesResponseDtos = new ArrayList<>();
    int index = 1;
    for (Machine machine : machines) {
        machinesResponseDtos.add(DaoToDto.convertMachineToEntityOptionResponseDto(index++, machine));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), machinesResponseDtos);
}
@GetMapping("/getAllMaintenanceRequests")
public AllEntityOptionResponseDto getAllMaintenanceRequests(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<MaintenanceRequest> maintenanceRequests = maintenanceRequestService.getAllMaintenanceRequests(companyId);
    List<EntityOptionResponseDto> maintenanceRequestsResponseDtos = new ArrayList<>();
    int index = 1;
    for (MaintenanceRequest maintenanceRequest : maintenanceRequests) {
        maintenanceRequestsResponseDtos.add(DaoToDto.convertMaintenanceRequestToEntityOptionResponseDto(index++, maintenanceRequest));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), maintenanceRequestsResponseDtos);
}
@GetMapping("/getAllActiveMaintenanceRequests")
public AllEntityOptionResponseDto getAllActiveMaintenanceRequests(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId) {
    List<MaintenanceRequest> maintenanceRequests = maintenanceRequestService.getAllActiveMaintenanceRequests(companyId);
    List<EntityOptionResponseDto> maintenanceRequestsResponseDtos = new ArrayList<>();
    int index = 1;
    for (MaintenanceRequest maintenanceRequest : maintenanceRequests) {
        maintenanceRequestsResponseDtos.add(DaoToDto.convertMaintenanceRequestToEntityOptionResponseDto(index++, maintenanceRequest));
    }
    return new AllEntityOptionResponseDto("Success", 200, LocalDateTime.now(), LocalDateTime.now(), maintenanceRequestsResponseDtos);
}
*/

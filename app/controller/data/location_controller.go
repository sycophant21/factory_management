package data

import (
	"factory_management_go/app/domain/dao/location"
	service "factory_management_go/app/service/data"
	httpUtil "factory_management_go/app/util/http"
	"net/http"
)

type LocationController struct {
	Mutex           *http.ServeMux
	LocationService *service.LocationService
}

func (lc *LocationController) Initialise() {
	lc.Mutex = http.NewServeMux()
	lc.Mutex.HandleFunc("GET /getAllLocationsFromLocationTypeId", lc.GetAllLocationsFromLocationTypeId)
	lc.Mutex.HandleFunc("GET /getAllLocationsFromLocationTypeCode", lc.GetAllLocationsFromLocationTypeCode)
	lc.Mutex.HandleFunc("GET /getAllLocations", lc.GetAllLocations)
	lc.Mutex.HandleFunc("GET /getLocationDetails", lc.GetLocationDetails)
	lc.Mutex.HandleFunc("GET /viewLocationDetails", lc.GetLocationDetails)
}

func (lc *LocationController) GetAllLocationsFromLocationTypeId(writer http.ResponseWriter, request *http.Request) {
	httpUtil.HandleRequest[[]*location.Location](writer, func() ([]*location.Location, error) {
		return lc.LocationService.GetAllLocationsFromLocationTypeId(request.URL.Query().Get("locationTypeId"), request.Header.Get("Company-Id"))
	}, func(data []*location.Location) interface{} {
		return httpUtil.ConvertAllLocationsToLocationResponseDto(data)
	})
}

func (lc *LocationController) GetAllLocationsFromLocationTypeCode(writer http.ResponseWriter, request *http.Request) {
	httpUtil.HandleRequest[[]*location.Location](writer, func() ([]*location.Location, error) {
		return lc.LocationService.GetAllLocationsFromLocationTypeCode(request.URL.Query().Get("locationTypeCode"), request.Header.Get("Company-Id"))
	}, func(data []*location.Location) interface{} {
		return httpUtil.ConvertAllLocationsToLocationResponseDto(data)
	})
}
func (lc *LocationController) GetAllLocations(writer http.ResponseWriter, request *http.Request) {
	httpUtil.HandleRequest[[]*location.Location](writer, func() ([]*location.Location, error) {
		return lc.LocationService.GetAllLocations(request.Header.Get("Company-Id"))
	}, func(data []*location.Location) interface{} {
		return httpUtil.ConvertAllLocationsToLocationResponseDto(data)
	})
}

func (lc *LocationController) GetLocationDetails(writer http.ResponseWriter, request *http.Request) {
	httpUtil.HandleRequest[location.Location](writer, func() (location.Location, error) {
		return lc.LocationService.GetLocationDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
	}, func(data location.Location) interface{} {
		return httpUtil.ConvertLocationToLocationResponseDtoForEdit(data)
	})
}
func (lc *LocationController) ViewLocationDetails(writer http.ResponseWriter, request *http.Request) {
	httpUtil.HandleRequest[location.Location](writer, func() (location.Location, error) {
		return lc.LocationService.GetLocationDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
	}, func(data location.Location) interface{} {
		return httpUtil.ConvertLocationToLocationResponseDtoForView(data)
	})
}

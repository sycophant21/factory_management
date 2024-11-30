package data

import (
	service "factory_management_go/app/service/data"
	"factory_management_go/app/util"
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
}

func (lc *LocationController) GetAllLocationsFromLocationTypeId(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return lc.LocationService.GetAllLocationsFromLocationTypeId(request.URL.Query().Get("locationTypeId"), request.Header.Get("Company-Id"))
	})
}

func (lc *LocationController) GetAllLocationsFromLocationTypeCode(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return lc.LocationService.GetAllLocationsFromLocationTypeCode(request.URL.Query().Get("locationTypeCode"), request.Header.Get("Company-Id"))
	})
}
func (lc *LocationController) GetAllLocations(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return lc.LocationService.GetAllLocations(request.Header.Get("Company-Id"))
	})
}

func (lc *LocationController) GetLocationDetails(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return lc.LocationService.GetLocationDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
	})
}

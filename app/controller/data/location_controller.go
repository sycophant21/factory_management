package data

import (
	service "factory_management_go/app/service/data"
	"factory_management_go/app/util"
	"net/http"
)

type LocationController struct {
	Mutex   *http.ServeMux
	Service *service.LocationService
}

func (l *LocationController) Initialise() {
	l.Mutex = http.NewServeMux()
	l.Mutex.HandleFunc("GET /getAllLocationsFromLocationType", l.GetAllLocationsFromLocationType)
	l.Mutex.HandleFunc("GET /getAllLocations", l.GetAllLocations)
	l.Mutex.HandleFunc("GET /getLocationDetails", l.GetLocationDetails)
}

func (l *LocationController) GetAllLocationsFromLocationType(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return l.Service.GetAllLocationsFromLocationType(request.URL.Query().Get("locationType"), request.Header.Get("Company-Id"))
	})
}
func (l *LocationController) GetAllLocations(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return l.Service.GetAllLocations(request.Header.Get("Company-Id"))
	})
}

func (l *LocationController) GetLocationDetails(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return l.Service.GetLocationDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
	})
}

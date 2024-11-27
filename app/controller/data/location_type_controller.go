package data

import (
	service "factory_management_go/app/service/data"
	"factory_management_go/app/util"
	"net/http"
)

type LocationTypeController struct {
	Mutex   *http.ServeMux
	Service *service.LocationTypeService
}

func (l *LocationTypeController) Initialise() {
	l.Mutex = http.NewServeMux()
	l.Mutex.HandleFunc("GET /getAllLocationTypes", l.GetAllLocationTypes)
	l.Mutex.HandleFunc("GET /getLocationTypeDetails", l.GetLocationTypeDetails)
	l.Mutex.HandleFunc("GET /getLocationTypeDetailsFromCode", l.GetLocationTypeDetailsFromCode)
}

func (l *LocationTypeController) GetAllLocationTypes(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return l.Service.GetAllLocationTypes(request.Header.Get("Company-Id"))
	})
}

func (l *LocationTypeController) GetLocationTypeDetails(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return l.Service.GetLocationTypeDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
	})
}

func (l *LocationTypeController) GetLocationTypeDetailsFromCode(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return l.Service.GetLocationTypeDetailsFromCode(request.URL.Query().Get("code"), request.Header.Get("Company-Id"))
	})
}

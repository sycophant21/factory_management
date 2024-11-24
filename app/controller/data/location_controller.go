package data

import (
	"encoding/json"
	service "factory_management_go/app/service/data"
	"fmt"
	"log"
	"net/http"
	"time"
)

type LocationController struct {
	LocationMutex *http.ServeMux
	Service       *service.LocationService
}

func (l *LocationController) Initialise() {
	l.LocationMutex = http.NewServeMux()
	l.LocationMutex.HandleFunc("GET /getAllLocationsFromLocationType", l.GetAllLocationsFromLocationType)
	l.Service = &service.LocationService{}
	l.Service.Initialise()
}

func (l *LocationController) GetAllLocationsFromLocationType(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/dataJson")
	locations, err := l.Service.GetAllLocationsFromLocationType(request.URL.Query().Get("locationType"), request.Header.Get("Company-Id"))
	var code = 200
	var message = "Success"
	var data = locations
	if err != nil {
		code = 500
		message = err.Error()
	}
	writer.WriteHeader(code)
	dataJson, err := json.Marshal(map[string]interface{}{"status": code, "data": data, "createdAt": time.Now(), "updatedAt": time.Now(), "message": message})
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprint(writer, string(dataJson))
	if err != nil {
		return
	}
}
func (l *LocationController) GetAllLocations(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/dataJson")
	locations, err := l.Service.GetAllLocations(request.Header.Get("Company-Id"))
	var code = 200
	var message = "Success"
	var data = locations
	if err != nil {
		code = 500
		message = err.Error()
	}
	writer.WriteHeader(code)
	dataJson, err := json.Marshal(map[string]interface{}{"status": code, "data": data, "createdAt": time.Now(), "updatedAt": time.Now(), "message": message})
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprint(writer, string(dataJson))
	if err != nil {
		return
	}
}

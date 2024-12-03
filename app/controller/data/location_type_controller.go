package data

import (
	"factory_management_go/app/domain/dao/location"
	service "factory_management_go/app/service/data"
	util "factory_management_go/app/util/http"
	"net/http"
)

type LocationTypeController struct {
	Mutex               *http.ServeMux
	LocationTypeService *service.LocationTypeService
}

func (ltc *LocationTypeController) Initialise() {
	ltc.Mutex = http.NewServeMux()
	ltc.Mutex.HandleFunc("GET /getAllLocationTypes", ltc.GetAllLocationTypes)
	ltc.Mutex.HandleFunc("GET /getLocationTypeDetails", ltc.GetLocationTypeDetails)
	ltc.Mutex.HandleFunc("GET /getLocationTypeDetailsFromCode", ltc.GetLocationTypeDetailsFromCode)
}

func (ltc *LocationTypeController) GetAllLocationTypes(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest[[]*location.LocationType](writer, func() ([]*location.LocationType, error) {
		return ltc.LocationTypeService.GetAllLocationTypes(request.Header.Get("Company-Id"))
	}, func(data []*location.LocationType) interface{} {
		return util.ConvertAllLocationTypesToLocationTypeResponseDto(data)
	})
}

func (ltc *LocationTypeController) GetLocationTypeDetails(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest[location.LocationType](writer, func() (location.LocationType, error) {
		return ltc.LocationTypeService.GetLocationTypeDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
	}, func(data location.LocationType) interface{} {
		return util.ConvertLocationTypeToLocationTypeResponseDto(data)
	})
}

func (ltc *LocationTypeController) GetLocationTypeDetailsFromCode(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest[location.LocationType](writer, func() (location.LocationType, error) {
		return ltc.LocationTypeService.GetLocationTypeDetailsFromCode(request.URL.Query().Get("code"), request.Header.Get("Company-Id"))
	}, func(data location.LocationType) interface{} {
		return util.ConvertLocationTypeToLocationTypeResponseDto(data)
	})
}

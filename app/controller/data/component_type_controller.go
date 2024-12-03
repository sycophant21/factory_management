package data

import (
	"factory_management_go/app/domain/dao/component"
	service "factory_management_go/app/service/data"
	util "factory_management_go/app/util/http"
	"net/http"
)

type ComponentTypeController struct {
	Mutex                *http.ServeMux
	ComponentTypeService *service.ComponentTypeService
}

func (ctc *ComponentTypeController) Initialise() {
	ctc.Mutex = http.NewServeMux()
	ctc.Mutex.HandleFunc("GET /getAllComponentTypes", ctc.GetAllComponentTypes)
	ctc.Mutex.HandleFunc("GET /getComponentType", ctc.GetComponentType)
}

func (ctc *ComponentTypeController) GetAllComponentTypes(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest[[]*component.ComponentType](writer,
		func() ([]*component.ComponentType, error) {
			return ctc.ComponentTypeService.GetAllComponentTypes(request.Header.Get("Company-Id"))
		},
		func(data []*component.ComponentType) interface{} {
			return util.ConvertAllComponentTypesToComponentTypeResponseDto(data)
		},
	)
}

func (ctc *ComponentTypeController) GetComponentType(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest[component.ComponentType](writer,
		func() (component.ComponentType, error) {
			return ctc.ComponentTypeService.GetComponentTypeDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
		},
		func(data component.ComponentType) interface{} {
			return util.ConvertComponentTypeToComponentTypeResponseDto(data)
		},
	)
}

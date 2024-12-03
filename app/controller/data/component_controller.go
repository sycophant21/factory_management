package data

import (
	"factory_management_go/app/domain/dao/component"
	service "factory_management_go/app/service/data"
	util "factory_management_go/app/util/http"
	"net/http"
)

type ComponentController struct {
	Mutex            *http.ServeMux
	ComponentService *service.ComponentService
}

func (cc *ComponentController) Initialise() {
	cc.Mutex = http.NewServeMux()
	cc.Mutex.HandleFunc("GET /getAllSpareTypes", cc.GetAllComponents)
	cc.Mutex.HandleFunc("GET /getAllComponents", cc.GetAllComponents)
	cc.Mutex.HandleFunc("GET /getSpareTypeDetails", cc.GetComponentDetails)
	cc.Mutex.HandleFunc("GET /getComponentDetails", cc.GetComponentDetails)
	cc.Mutex.HandleFunc("GET /viewSpareTypeDetails", cc.ViewComponentDetails)
	cc.Mutex.HandleFunc("GET /viewComponentDetails", cc.ViewComponentDetails)

}

func (cc *ComponentController) GetAllComponents(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest[[]*component.Component](writer,
		func() ([]*component.Component, error) {
			return cc.ComponentService.GetAllComponents(request.Header.Get("Company-Id"))
		},
		func(data []*component.Component) interface{} {
			return util.ConvertAllComponentsToComponentResponseDto(data)
		},
	)
}

func (cc *ComponentController) GetComponentDetails(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest[component.Component](writer,
		func() (component.Component, error) {
			return cc.ComponentService.GetComponentDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
		},
		func(data component.Component) interface{} {
			return util.ConvertComponentToComponentResponseDtoForEdit(data)
		},
	)
}

func (cc *ComponentController) ViewComponentDetails(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest[component.Component](writer,
		func() (component.Component, error) {
			return cc.ComponentService.GetComponentDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
		},
		func(data component.Component) interface{} {
			return util.ConvertComponentToComponentResponseDtoForView(data)
		},
	)
}

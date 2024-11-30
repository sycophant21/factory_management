package data

import (
	service "factory_management_go/app/service/data"
	"net/http"
)

type ComponentTypeController struct {
	Mutex                *http.ServeMux
	ComponentTypeService *service.ComponentTypeService
}

func (ctc *ComponentTypeController) Initialise() {
	ctc.Mutex = http.NewServeMux()
}

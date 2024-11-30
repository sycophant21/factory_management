package data

import (
	service "factory_management_go/app/service/data"
	util "factory_management_go/app/util"
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
}
func (cc *ComponentController) GetAllComponents(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return cc.ComponentService.GetAllComponents(request.Header.Get("Company-Id"))
	})
}

func (cc *ComponentController) GetComponentDetails(writer http.ResponseWriter, request *http.Request) {
	util.HandleRequest(writer, func() (interface{}, error) {
		return cc.ComponentService.GetComponentDetails(request.URL.Query().Get("id"), request.Header.Get("Company-Id"))
	})
}

/*
   @GetMapping(value = {"/getSpareTypeDetails", "/getComponentDetails"})
   public ComponentResponseDto getComponentDetails(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId, @NotNull @RequestParam("id") String id) {
       try {
           return DaoToDto.convertComponentToComponentResponseDtoForEdit(componentService.getComponent(companyId, id));
       } catch (InvalidInputException iie) {
           return new ComponentResponseDto("Failure", 400, LocalDateTime.now(), LocalDateTime.now(), "", "", "", "", "", "", "", "", "", -1, -1, "", "", "");
       }
   }

   @GetMapping(value = {"/viewSpareTypeDetails", "/viewComponentDetails"})
   public ComponentResponseDto viewComponentDetails(@NotNull @RequestHeader("user-id") String userId, @NotNull @RequestHeader("company-id") String companyId, @NotNull @RequestParam("id") String id) {
       try {
           return DaoToDto.convertComponentToComponentResponseDtoForView(componentService.getComponent(companyId, id));
       } catch (InvalidInputException iie) {
           return new ComponentResponseDto("Failure", 400, LocalDateTime.now(), LocalDateTime.now(), "", "", "", "", "", "", "", "", "", -1, -1, "", "", "");
       }
   }
*/

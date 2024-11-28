package component

import (
	"factory_management_go/app/domain/dto"
	"factory_management_go/app/domain/dto/request/wrapper"
	"time"
)

type UpsertComponentRequestDto struct {
	wrapper.RequestDto
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	ComponentType        string `json:"componentType"`
	ComponentUseCategory string `json:"componentUseCategory"`
	ComponentReturnType  string `json:"componentReturnType"`
	ComponentPosition    string `json:"componentPosition"`
	ExternalUnit         string `json:"conversionFactor_externalUnitData_externalUnit"`
	ExternalUnits        uint16 `json:"conversionFactor_externalUnitData_externalUnitQuantity"`
	InternalUnit         string `json:"conversionFactor_internalUnitData_internalUnit"`
	InternalUnits        uint16 `json:"conversionFactor_internalUnitData_internalUnitQuantity"`
	HsnCode              string `json:"hsnCode"`
	GstPercentage        string `json:"gstPercentage"`
}

type UpsertComponentInventoryRequestDto struct {
	wrapper.RequestDto
	Id                        string `json:"id"`
	Name                      string `json:"name"`
	PrintName                 string `json:"printName"`
	Description               string `json:"description"`
	Note                      string `json:"note"`
	Component                 string `json:"component_id"`
	ComponentUseCategory      string `json:"componentUseCategory"`
	ComponentReturnType       string `json:"componentReturnType"`
	ComponentPosition         string `json:"componentPosition"`
	ExternalUnit              string `json:"conversionFactor_externalUnitData_externalUnit"`
	ExternalUnits             uint16 `json:"conversionFactor_externalUnitData_externalUnitQuantity"`
	InternalUnit              string `json:"conversionFactor_internalUnitData_internalUnit"`
	InternalUnits             uint16 `json:"conversionFactor_internalUnitData_internalUnitQuantity"`
	HsnCode                   string `json:"hsnCode"`
	GstPercentage             string `json:"gstPercentage"`
	Location                  string `json:"location"`
	MaintainBatches           bool   `json:"maintainBatches"`
	MinStock                  uint16 `json:"minStock"`
	MaxStock                  uint16 `json:"maxStock"`
	CurrentStock              uint16 `json:"currentStock"`
	HaveWarranty              bool   `json:"haveWarranty"`
	ReorderQuantity           uint16 `json:"reorderQuantity"`
	StandardPrice             uint16 `json:"standardPrice"`
	Active                    bool   `json:"active"`
	ComponentLifeDurationType string `json:"life_componentLifeDurationType"`
	ComponentLifeDuration     uint16 `json:"life_componentLifeDuration"`
}

type UpsertLocationRequestDto struct {
	wrapper.RequestDto
	Id           string `json:"id"`
	Name         string `json:"name"`
	Details      string `json:"details"`
	LocationType string `json:"locationFor"`
}

type UpsertMachineRequestDto struct {
	wrapper.RequestDto
	Id                          string     `json:"id"`
	Name                        string     `json:"name"`
	MachineIndex                string     `json:"machineIndex"`
	Description                 string     `json:"description"`
	MachineType                 string     `json:"machineType"`
	MachineStatus               string     `json:"machineStatus"`
	Movable                     bool       `json:"movable"`
	MachineUsage                string     `json:"machineUsage"`
	Location                    string     `json:"location"`
	Department                  string     `json:"department"`
	Supplier                    string     `json:"supplier"`
	Manufacturer                string     `json:"manufacturer"`
	BillNumber                  string     `json:"billNumber"`
	SerialNumber                string     `json:"serialNumber"`
	ModelNumber                 string     `json:"modelNumber"`
	HsnCode                     string     `json:"hsnCode"`
	GstPercentage               string     `json:"gstPercentage"`
	Warranty                    bool       `json:"warranty"`
	WarrantyInfo                string     `json:"warrantyInfo"`
	WarrantyExpirationDate      *time.Time `json:"warrantyExpirationDate"`
	PeriodicMaintenanceRequired bool       `json:"periodicMaintenanceRequired"`
	FirstMajorServiceDate       *time.Time `json:"machineServiceDateInformation_machineMajorServiceDate_machineFirstMajorServiceDate"`
	MajorServiceInterval        uint16     `json:"machineServiceDateInformation_machineMajorServiceDate_machineMajorServiceDateInterval"`
	MajorServiceIntervalType    string     `json:"machineServiceDateInformation_machineMajorServiceDate_machineMajorServiceDateIntervalType"`
	FirstMinorServiceDate       *time.Time `json:"machineServiceDateInformation_machineMinorServiceDate_machineFirstMinorServiceDate"`
	MinorServiceInterval        uint16     `json:"machineServiceDateInformation_machineMinorServiceDate_machineMinorServiceDateInterval"`
	MinorServiceIntervalType    string     `json:"machineServiceDateInformation_machineMinorServiceDate_machineMinorServiceDateIntervalType"`
	Components                  []string   `json:"components"`
}

type UpsertMaintenanceRequestDto struct {
	wrapper.RequestDto
	Id                     string                  `json:"id"`
	MachineId              string                  `json:"machineId"`
	Details                string                  `json:"details"`
	ComponentsRequired     map[string]dto.Quantity `json:"componentsRequired"`
	ScheduledStartDateTime *time.Time              `json:"scheduledStartDateTime"`
	ScheduledEndDateTime   *time.Time              `json:"scheduledEndDateTime"`
	MaintenanceBy          string                  `json:"maintenanceBy"`
	MaintenanceTeamNames   string                  `json:"maintenanceTeamNames"`
	ServiceVendor          string                  `json:"serviceVendor"`
}

type UpsertPurchaseRequestDto struct {
	wrapper.RequestDto
	Id             string                  `json:"id"`
	Details        string                  `json:"details"`
	PurchaseDate   *time.Time              `json:"purchasedDate"`
	Vendor         string                  `json:"vendor"`
	ItemsPurchased map[string]dto.Quantity `json:"itemsPurchased"`
}

type UpsertServiceVendorRequestDto struct {
	wrapper.RequestDto
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	PhoneNumber string   `json:"phoneNumber"`
	Address     string   `json:"address"`
	Machines    []string `json:"machines"`
}

type UpsertSparePartVendorRequestDto struct {
	wrapper.RequestDto
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	PhoneNumber string   `json:"phoneNumber"`
	Address     string   `json:"address"`
	SpareParts  []string `json:"spareParts"`
}

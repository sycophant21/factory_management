package machine

import (
	"factory_management_go/app/domain/dao"
	"factory_management_go/app/domain/dao/location"
	"factory_management_go/app/domain/dao/user"
	"factory_management_go/app/domain/metadata"
	"time"
)

type (
	Machine struct {
		Id                          string               `xorm:"'id' pk"`
		MachineIndex                uint8                `xorm:"'machine_index'"`
		Description                 string               `xorm:"'description'"`
		LocationId                  string               `xorm:"'location_id'"`
		Location                    location.Location    `xorm:"-"`
		DepartmentId                string               `xorm:"'department_id'"`
		Department                  dao.Department       `xorm:"-"`
		SupplierId                  string               `xorm:"'supplier_id'"`
		Supplier                    user.Supplier        `xorm:"-"`
		ManufacturerId              string               `xorm:"'manufacturer_id'"`
		Manufacturer                user.Manufacturer    `xorm:"-"`
		MachineType                 string               `xorm:"'machine_type'"`
		MachineStatusId             string               `xorm:"'machine_status_id'"`
		MachineStatus               MachineStatus        `xorm:"-"`
		Movable                     bool                 `xorm:"'movable'"`
		MachineUsageId              string               `xorm:"'machine_usage_id'"`
		MachineUsage                MachineUsage         `xorm:"-"`
		BillNumber                  string               `xorm:"'bill_number'"`
		SerialNumber                string               `xorm:"'serial_number'"`
		ModelNumber                 string               `xorm:"'model_number'"`
		HsnCode                     string               `xorm:"'hsn_code'"`
		GstPercentageId             string               `xorm:"'gst_percentage_id'"`
		GSTPercentage               dao.GSTPercentage    `xorm:"-"`
		Warranty                    bool                 `xorm:"'warranty'"`
		WarrantyInfo                string               `xorm:"'warranty_info'"`
		WarrantyExpirationDate      *time.Time           `xorm:"warranty_expiration_date'"`
		PeriodicMaintenanceRequired bool                 `xorm:"'periodic_maintenance_required'"`
		MajorServiceDateInfo        dao.PeriodicDateInfo `xorm:"extends"`
		MinorServiceDateInfo        dao.PeriodicDateInfo `xorm:"extends"`
		Metadata                    metadata.Metadata    `xorm:"extends"`
	}
	MachineStatus struct {
		Id       string            `xorm:"'id' pk"`
		Name     string            `xorm:"status_name"`
		Code     string            `xorm:"status_code"`
		metadata metadata.Metadata `xorm:"extends"`
	}
	MachineUsage struct {
		Id       string            `xorm:"'id' pk"`
		name     string            `xorm:"usage_type_name"`
		code     string            `xorm:"usage_type_code"`
		metadata metadata.Metadata `xorm:"extends"`
	}
	//MachineType unused
	MachineType struct{}
	//MachineComponent unused
	MachineComponent struct{}
)

func (ms *MachineStatus) TableName() string {
	return "enum_machine_status"
}

func (mu *MachineUsage) TableName() string {
	return "enum_machine_usage"
}

/*
   @ManyToMany
   private List<ComponentInventory> machineComponents;
*/

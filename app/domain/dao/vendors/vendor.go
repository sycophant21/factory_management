package vendors

import (
	"factory_management_go/app/domain/dao/component"
	"factory_management_go/app/domain/dao/machine"
	"factory_management_go/app/domain/metadata"
)

type (
	ServiceVendor struct {
		Id          string               `xorm:"'id' pk"`
		Name        string               `xorm:"'name'"`
		PhoneNumber metadata.PhoneNumber `xorm:"extends 'phone_number'"`
		Address     string               `xorm:"'address'"`
		Services    []machine.Machine    `xorm:"-"`
		Metadata    metadata.Metadata    `xorm:"extends"`
	}

	ServiceVendorServices struct {
		ServicesId      string `xorm:"'services_id'"`
		ServiceVendorId string `xorm:"'service_vendor_id'"`
	}

	SparePartVendor struct {
		Id           string                `xorm:"'id' pk"`
		Name         string                `xorm:"'name'"`
		PhoneNumber  metadata.PhoneNumber  `xorm:"extends 'phone_number'"`
		Address      string                `xorm:"'address'"`
		SoldByVendor []component.SparePart `xorm:"-"`
		Metadata     metadata.Metadata     `xorm:"extends"`
	}

	SparePartVendorSoldByVendor struct {
		SoldByVendorId    string `xorm:"'sold_by_vendor_id'"`
		SparePartVendorId string `xorm:"'spare_part_vendor_id'"`
	}
)

func (sv *ServiceVendor) TableName() string {
	return "service_vendor"
}
func (svs *ServiceVendorServices) TableName() string {
	return "service_vendor_services"
}

func (sv *SparePartVendor) TableName() string {
	return "spare_part_vendor"
}

func (sv *SparePartVendorSoldByVendor) TableName() string {
	return "spare_part_vendor_sold_by_vendor"
}

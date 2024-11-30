package inventory

import (
	"factory_management_go/app/domain/dao/component"
	"factory_management_go/app/domain/dao/vendors"
	"factory_management_go/app/domain/metadata"
	"time"
)

type (
	Purchase struct {
		Id           string                                   `xorm:"'id' pk"`
		Details      string                                   `xorm:"'details'"`
		PurchaseDate *time.Time                               `xorm:"'purchase_date'"`
		VendorId     string                                   `xorm:"'vendor_id'"`
		Vendor       *vendors.SparePartVendor                 `xorm:"-"`
		Items        map[*component.ComponentInventory]uint16 `xorm:"-"`
		Metadata     *metadata.Metadata                       `xorm:"extends"`
	}
	PurchaseItems struct {
		Items          uint16 `xorm:"'items'"`
		PartPurchaseId string `xorm:"'items_key'"`
		PurchaseId     string `xorm:"'purchase_id'"`
	}
)

func (p *Purchase) TableName() string {
	return "purchase"
}
func (pi *PurchaseItems) TableName() string {
	return "purchase_items"
}

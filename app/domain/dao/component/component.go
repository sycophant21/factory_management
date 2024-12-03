package component

import (
	"factory_management_go/app/domain/dao"
	"factory_management_go/app/domain/dao/location"
	"factory_management_go/app/domain/metadata"
)

type (
	Component struct {
		Id              string             `xorm:"pk 'id'"`
		Name            string             `xorm:"'name'"`
		Description     string             `xorm:"'description'"`
		ComponentTypeId string             `xorm:"'component_type_id'"`
		ComponentType   *ComponentType     `xorm:"-"`
		ComponentInfo   *ComponentInfo     `xorm:"extends"`
		Metadata        *metadata.Metadata `xorm:"extends"`
	}
	SpareType struct {
		Id              string             `xorm:"pk 'id'"`
		Name            string             `xorm:"'name'"`
		Description     string             `xorm:"'description'"`
		ComponentTypeId string             `xorm:"'component_type_id'"`
		ComponentType   *ComponentType     `xorm:"-"`
		ComponentInfo   *ComponentInfo     `xorm:"extends"`
		Metadata        *metadata.Metadata `xorm:"extends"`
	}
	ComponentInfo struct {
		ComponentUseCategoryId string                `xorm:"'component_use_category_id'"`
		ComponentUseCategory   *ComponentUseCategory `xorm:"-"`
		OldPartReturnTypeId    string                `xorm:"'old_part_return_type_id'"`
		OldPartReturnType      *ComponentReturnType  `xorm:"-"`
		OldPartPositionId      string                `xorm:"'old_part_position_id'"`
		OldPartPosition        *ComponentPosition    `xorm:"-"`
		ExternalUnitId         string                `xorm:"'external_unit_id'"`
		ExternalUnit           *ComponentUnit        `xorm:"-"`
		InternalUnitId         string                `xorm:"'internal_unit_id'"`
		InternalUnit           *ComponentUnit        `xorm:"-"`
		ConversionFactor       *dao.ConversionFactor `xorm:"extends"`
		HsnCode                string                `xorm:"'hsn_code'"`
		GstPercentageId        string                `xorm:"'gst_percentage_id'"`
		GstPercentage          *dao.GSTPercentage    `xorm:"-"`
	}
	ComponentInventory struct {
		Id                        string                     `xorm:"'id'"`
		Name                      string                     `xorm:"'name'"`
		PrintName                 string                     `xorm:"'print_name'"`
		Description               string                     `xorm:"'description'"`
		Note                      string                     `xorm:"'note'"`
		ComponentId               string                     `xorm:"'component_id'"`
		Component                 *Component                 `xorm:"-"`
		ComponentInventoryInfo    *ComponentInfo             `xorm:"extends"`
		ComponentStockInformation *ComponentStockInformation `xorm:"extends"`
		Metadata                  *metadata.Metadata         `xorm:"extends"`
	}
	SparePart struct {
		Id                        string                     `xorm:"'id'"`
		Name                      string                     `xorm:"'name'"`
		PrintName                 string                     `xorm:"'print_name'"`
		Description               string                     `xorm:"'description'"`
		Note                      string                     `xorm:"'note'"`
		ComponentId               string                     `xorm:"'component_id'"`
		Component                 *Component                 `xorm:"-"`
		ComponentInfo             *ComponentInfo             `xorm:"extends"`
		ComponentStockInformation *ComponentStockInformation `xorm:"extends"`
		Metadata                  *metadata.Metadata         `xorm:"extends"`
	}
	ComponentLife struct {
		DurationTypeId string            `xorm:"'duration_type_id'"`
		DurationType   *dao.DurationType `xorm:"-"`
		Duration       uint8             `xorm:"'duration'"`
	}
	ComponentPosition struct {
		Id       string             `xorm:"'id'"`
		Name     string             `xorm:"'position_name'"`
		Code     string             `xorm:"'position_code'"`
		Metadata *metadata.Metadata `xorm:"extends"`
	}
	ComponentReturnType struct {
		Id       string             `xorm:"'id'"`
		Name     string             `xorm:"'return_type_name'"`
		Code     string             `xorm:"'return_type_code'"`
		Metadata *metadata.Metadata `xorm:"extends"`
	}
	ComponentStockInformation struct {
		LocationId      string             `xorm:"'location_id'"`
		Location        *location.Location `xorm:"-"`
		MaintainBatches bool               `xorm:"'maintain_batches'"`
		MinStock        uint32             `xorm:"'min_stock'"`
		MaxStock        uint32             `xorm:"'max_stock'"`
		CurrentStock    uint32             `xorm:"'current_stock'"`
		HaveWarranty    bool               `xorm:"'have_warranty'"`
		ReorderQuantity uint32             `xorm:"'reorder_quantity'"`
		StandardPrice   uint32             `xorm:"'standard_price'"`
		Active          bool               `xorm:"'active'"`
		Life            *ComponentLife     `xorm:"extends"`
		Barcode         string             `xorm:"'barcode'"`
		Image           []byte             `xorm:"'image'"`
	}
	ComponentType struct {
		Id       string             `xorm:"'id'"`
		Name     string             `xorm:"'name'"`
		Code     string             `xorm:"'code'"`
		Metadata *metadata.Metadata `xorm:"extends"`
	}
	ComponentUnit struct {
		Id       string             `xorm:"'id'"`
		Name     string             `xorm:"'name'"`
		Code     string             `xorm:"'code'"`
		Metadata *metadata.Metadata `xorm:"extends"`
	}
	ComponentUseCategory struct {
		Id       string             `xorm:"'id'"`
		Name     string             `xorm:"'category_name'"`
		Code     string             `xorm:"'category_code'"`
		Metadata *metadata.Metadata `xorm:"extends"`
	}
)

// Override Methods from "Unknown" Interface (in #xorm) (TableName),
// Needed if custom table name is required, Added to all to be on the safer side
func (c *Component) TableName() string {
	return "component"
}

// Override Methods from "Unknown" Interface (in #xorm) (TableName),
// Needed if custom table name is required, Added to all to be on the safer side
func (st *SpareType) TableName() string {
	return "component"
}

// Override Methods from "Unknown" Interface (in #xorm) (TableName),
// Needed if custom table name is required, Added to all to be on the safer side
func (sp *SparePart) TableName() string {
	return "component_inventory"
}

// Override Methods from "Unknown" Interface (in #xorm) (TableName),
// Needed if custom table name is required, Added to all to be on the safer side
func (ci *ComponentInventory) TableName() string {
	return "component_inventory"
}

// Override Methods from IOptions Interface (in #app/domain/option/options.go) (GetIndex, GetLabel, GetValue)
func (ct *ComponentType) GetLabel() string {
	return ct.Name
}
func (ct *ComponentType) GetValue() string {
	return ct.Code
}

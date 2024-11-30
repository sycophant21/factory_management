package location

import (
	"factory_management_go/app/domain/metadata"
)

type (
	Location struct {
		Id             string             `xorm:"pk"`
		Name           string             `xorm:"name"`
		Details        string             `xorm:"details"`
		LocationTypeId string             `xorm:"location_type_id"`
		LocationType   *LocationType      `xorm:"-"`
		Metadata       *metadata.Metadata `xorm:"extends"`
	}
	LocationType struct {
		Id       string             `xorm:"pk"`
		Name     string             `xorm:"name"`
		Code     string             `xorm:"code"`
		Metadata *metadata.Metadata `xorm:"extends"`
	}
)

func (loc *Location) SetLocationType(lt *LocationType) {
	if loc.LocationType == nil {
		loc.LocationType = &LocationType{}
	}
	*loc.LocationType = *lt
}

// Override Methods from "Unknown" Interface (in xorm) (TableName),
// Needed if custom table name is required, Added to all to be on the safer side
func (loc *Location) TableName() string {
	return "location"
}

// Override Methods from IOptions Interface (in #app/domain/option/options.go) (GetIndex, GetLabel, GetValue)
func (loc *Location) GetLabel() string {
	return loc.Name
}
func (loc *Location) GetValue() string {
	return loc.Id
}

// Override Methods from "Unknown" Interface (in #xorm) (TableName),
// Needed if custom table name is required, Added to all to be on the safer side
func (lt *LocationType) TableName() string {
	return "location_type"
}

// Override Methods from IOptions Interface (in #app/domain/option/options.go) (GetIndex, GetLabel, GetValue)
func (lt *LocationType) GetLabel() string {
	return lt.Name
}
func (lt *LocationType) GetValue() string {
	return lt.Code
}

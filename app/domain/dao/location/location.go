package location

import (
	"factory_management_go/app/domain/metadata"
)

type (
	Location struct {
		Id             string            `xorm:"pk"`
		Name           string            `xorm:"name"`
		Details        string            `xorm:"details"`
		LocationTypeId string            `xorm:"location_type_id"`
		LocationType   *LocationType     `xorm:"-"`
		Metadata       metadata.Metadata `xorm:"extends"`
	}
	LocationType struct {
		Id       string            `xorm:"pk"`
		Name     string            `xorm:"name"`
		Code     string            `xorm:"code"`
		Metadata metadata.Metadata `xorm:"extends"`
	}
)

func (loc *Location) SetLocationType(lt *LocationType) {
	if loc.LocationType == nil {
		loc.LocationType = &LocationType{}
	}
	*loc.LocationType = *lt
}

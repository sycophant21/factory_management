package location

import (
	metadata "factory_management_go/app/domain/metadata"
)

type Location struct {
	Id             string            `xorm:"pk"`
	Name           string            `xorm:"name"`
	Details        string            `xorm:"details"`
	LocationTypeId string            `xorm:"location_type_id"`
	Metadata       metadata.Metadata `xorm:"extends"`
}

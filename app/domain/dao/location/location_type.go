package location

import (
	metadata "factory_management_go/app/domain/metadata"
)

type LocationType struct {
	Id       string `xorm:"pk"`
	Name     string
	Code     string
	Metadata metadata.Metadata `xorm:"extends"`
}

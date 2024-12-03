package dao

import "factory_management_go/app/domain/metadata"

type GSTPercentage struct {
	Id       string            `xorm:"'id'"`
	Value    float32           `xorm:"'value'"`
	Metadata metadata.Metadata `xorm:"extends"`
}

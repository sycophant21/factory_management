package dao

import "factory_management_go/app/domain/metadata"

type DurationType struct {
	Id       string             `xorm:"pk"`
	Name     string             `xorm:"periodic_date_duration_type_value"`
	Code     string             `xorm:"periodic_date_duration_type_code"`
	Metadata *metadata.Metadata `xorm:"extends"`
}

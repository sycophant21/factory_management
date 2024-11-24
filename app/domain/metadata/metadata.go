package metadata

import (
	"time"
)

type Metadata struct {
	CreationTimestamp    time.Time `xorm:"created"`
	LastUpdatedTimestamp time.Time `xorm:"updated"`
	CreatedById          string
	LastUpdatedById      string
	CompanyId            string
	IsActive             bool
}

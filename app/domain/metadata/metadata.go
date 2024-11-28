package metadata

import (
	"time"
)

type Metadata struct {
	CreationTimestamp    *time.Time `xorm:"created"`
	LastUpdatedTimestamp *time.Time `xorm:"updated"`
	CreatedById          string     `xorm:"'created_by_id'"`
	LastUpdatedById      string     `xorm:"'last_updated_by_id'"`
	CompanyId            string     `xorm:"'company_id'"`
	IsActive             bool       `xorm:"'is_active'"`
}

type Company struct {
	Id               string      `xorm:"'id' pk"`
	Name             string      `xorm:"'name'"`
	PublicIdentifier string      `xorm:"'public_identifier'"`
	PhoneNumber      PhoneNumber `xorm:"extends"`
}

type EmailAddress struct {
	EmailAddress string `xorm:"'email_address'"`
}

type PhoneNumber struct {
	PhoneNumber string `xorm:"'phone_number'"`
}

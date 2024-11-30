package dao

type ConversionFactor struct {
	ExternalUnits uint16 `xorm:"'external_units'"`
	InternalUnits uint16 `xorm:"'internal_units'"`
}

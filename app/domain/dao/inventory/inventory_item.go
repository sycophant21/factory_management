package inventory

import (
	"factory_management_go/app/domain/dao/component"
	"factory_management_go/app/domain/metadata"
)

type (
	InventoryItem struct {
		Id         string                          `xorm:"'id' pk"`
		ItemId     string                          `xorm:"'item_id'"`
		Item       component.SparePart             `xorm:"-"`
		Quantities map[*InventoryItemStatus]uint16 `xorm:"-"`
		Metadata   metadata.Metadata               `xorm:"extends"`
	}
	InventoryItemQuantities struct {
		Quantities          uint16              `xorm:"'quantities'"`
		InventoryItemStatus InventoryItemStatus `xorm:"'quantities_key'"`
		InventoryItemId     string              `xorm:"'inventory_item_id'"`
	}
	InventoryItemRequisition struct {
		Id                        string                     `xorm:"'id' pk"`
		InventoryItemTransactions []InventoryItemTransaction `xorm:"-"`
		Metadata                  metadata.Metadata          `xorm:"extends"`
	}
	InventoryItemRequisitionInventoryItemTransactions struct {
		InventoryItemRequisitionId string `xorm:"'inventory_item_requisition_id'"`
		InventoryItemTransactionId string `xorm:"'inventory_item_transactions_id'"`
	}
	InventoryItemStatus      string
	InventoryItemTransaction struct {
		Id              string                       `xorm:"'id' pk"`
		ItemId          string                       `xorm:"'item_id'"`
		Item            InventoryItem                `xorm:"_"`
		Quantity        uint16                       `xorm:"'quantity'"`
		TransactionType InventoryItemTransactionType `xorm:"'transactionType'"`
		Metadata        metadata.Metadata            `xorm:"extends"`
	}
	InventoryItemTransactionType string
)

const (
	IN_USE       InventoryItemStatus = "IN_USE"
	ON_HOLD      InventoryItemStatus = "ON_HOLD"
	IN_INVENTORY InventoryItemStatus = "IN_INVENTORY"

	REQUEST  InventoryItemTransactionType = "REQUEST"
	ALLOCATE InventoryItemTransactionType = "ALLOCATE"
	RETURN   InventoryItemTransactionType = "RETURN"
)

// Override Methods from IConversion Interface (FromDB, ToDB), Needed for Custom Types (Enums)
func (iis *InventoryItemStatus) FromDB(data []byte) error {
	if data == nil {
		*iis = InventoryItemStatus("")
		return nil
	}
	*iis = InventoryItemStatus(string(data))
	return nil
}
func (iis *InventoryItemStatus) ToDB() ([]byte, error) {
	if iis == nil {
		return nil, nil
	}
	return []byte(string(*iis)), nil
}

// Override Methods from IConversion Interface (FromDB, ToDB), Needed for Custom Types (Enums)
func (iitt *InventoryItemTransactionType) FromDB(data []byte) error {
	if data == nil {
		*iitt = InventoryItemTransactionType("")
		return nil
	}
	*iitt = InventoryItemTransactionType(string(data))
	return nil
}
func (iitt *InventoryItemTransactionType) ToDB() ([]byte, error) {
	if iitt == nil {
		return nil, nil
	}
	return []byte(string(*iitt)), nil
}

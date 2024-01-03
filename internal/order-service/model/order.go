package model

import "common"

type Order struct {
	common.SqlModel
	UserID     int              `json:"user_id" gorm:"column:user_id;not null;index"`
	MenuItemID string           `json:"menu_item_id" gorm:"column:menu_item_id;not null"`
	TotalBill  float32          `json:"total_bill" gorm:"column:total_bill;not null"`
	DiscountID int              `json:"discount_id" gorm:"column:discount_id"`
	Status     string           `json:"status" gorm:"column:status;not null"`
	Location   *common.Location `json:"location" gorm:"column:location;not null;type:json"`
}

const (
	OrderStatusPending    = "pending"
	OrderStatusPreparing  = "preparing"
	OrderStatusFinished   = "finished"
	OrderStatusDelivering = "delivering"
	OrderStatusDelivered  = "delivered"
)

func (o *Order) ChangeStatusToPending() {
	o.Status = OrderStatusPending
}

func (o *Order) ChangeStatusToPreparing() {
	o.Status = OrderStatusPreparing
}

func (o *Order) ChangeStatusToFinished() {
	o.Status = OrderStatusFinished
}

func (o *Order) ChangeStatusToDelivering() {
	o.Status = OrderStatusDelivering
}

func (o *Order) ChangeStatusToDelivered() {
	o.Status = OrderStatusDelivered
}

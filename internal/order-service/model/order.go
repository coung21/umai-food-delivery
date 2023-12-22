package model

import "common"

type Order struct {
	common.SqlModel
	UserID     int
	MenuItemID string
	TotalBill  float32
	Promo      int
	Status     string
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

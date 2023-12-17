package usecase

import (
	order "order-service/interfaces"
)

type orderUC struct {
	OrderRepo order.CacheRepository
}

func NewOrderUC(orderRepo order.CacheRepository) *orderUC {
	return &orderUC{OrderRepo: orderRepo}
}

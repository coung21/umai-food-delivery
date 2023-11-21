package common

import "errors"

const (
	CuserId   = "current_user_id"
	CuserRole = "current_user_role"
	CResId    = "current_restaurant_id"
)

const (
	RoleCustomer   = "customer"
	RoleRestaurant = "restaurant"
	RoleShipper    = "shipper"
	RoleAdmin      = "admin"
)

var (
	ErrMissCache = errors.New("Miss cache")
)

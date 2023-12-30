package model

type CartItem struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type DeletingReq struct {
	ItemsID []int `json:"items_id"`
}

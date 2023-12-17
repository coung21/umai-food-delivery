package model

type CartItem struct {
	ItemID   string `json:"item_id"`
	Quantity int    `json:"quantity"`
}

type DeletingReq struct {
	ItemsID []string `json:"items_id"`
}

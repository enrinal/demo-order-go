package entity

type Order struct {
	ID         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	Products   []Product `json:"products"`
	Status     string    `json:"status"`
	TotalPrice float64   `json:"total_price"`
	TotalQty   int64     `json:"total_qty"`
	TotalItem  int64     `json:"total_item"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
}

type OrderRequest struct {
	UserId int64 `json:"user_id"`
	CartId int64 `json:"cart_id"`
}

type ListOrdersResponse struct {
	Orders []Order `json:"orders"`
}

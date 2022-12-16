package entity

type Cart struct {
	ID        int64     `json:"id"`
	UserId    int64     `json:"user_id"`
	Products  []Product `json:"products"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type CartRequest struct {
	UserId   int64    `json:"user_id"`
	Products []string `json:"products"`
}

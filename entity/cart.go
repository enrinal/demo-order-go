package entity

import "errors"

type Cart struct {
	ID        string                `json:"id"`
	UserId    string                `json:"user_id"`
	Products  []CartProductResponse `json:"products"`
	CreatedAt string                `json:"created_at"`
	UpdatedAt string                `json:"updated_at"`
}

type CartProductRequest struct {
	Id  string `json:"id"`
	Qty int64  `json:"qty"`
}

type CartRequest struct {
	UserId   string               `json:"user_id"`
	Products []CartProductRequest `json:"products"`
}

type CartProductResponse struct {
	Product
	Qty int64 `json:"qty"`
}

func (cr *CartRequest) Validate() error {
	if len(cr.Products) == 0 {
		return errors.New("products is required")
	}

	if len(cr.Products) > 10 {
		return errors.New("maximum products is 10")
	}

	return nil
}

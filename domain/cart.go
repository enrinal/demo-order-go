package domain

import (
	"context"

	"github.com/enrinal/demo-order-go/entity"

	"github.com/enrinal/demo-order-go/models"
)

// for database

type CartRepository interface {
	Store(ctx context.Context, cart models.Cart) error
	FindById(ctx context.Context, id string) (*models.Cart, error)
}

// for business logic

type CartService interface {
	AddCart(ctx context.Context, req entity.CartRequest) error
	GetCartById(ctx context.Context, id string) (*entity.Cart, error)
}

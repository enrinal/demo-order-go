package domain

import (
	"context"

	"github.com/enrinal/demo-order-go/entity"

	"github.com/enrinal/demo-order-go/models"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]models.Product, error)
	FindById(ctx context.Context, id string) (*models.Product, error)
}

type ProductService interface {
	GetAll(ctx context.Context) ([]entity.Product, error)
	GetById(ctx context.Context, id string) (*entity.Product, error)
}

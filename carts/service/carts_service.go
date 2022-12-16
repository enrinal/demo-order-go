package service

import (
	"context"

	"github.com/enrinal/demo-order-go/models"

	"github.com/enrinal/demo-order-go/domain"
	"github.com/enrinal/demo-order-go/entity"
	"github.com/go-redis/redis/v8"
)

type cartService struct {
	cartRepo    domain.CartRepository
	productRepo domain.ProductRepository
	rc          *redis.Client
}

func NewCartService(cartRepo domain.CartRepository, productRepo domain.ProductRepository, rc *redis.Client) domain.CartService {
	return &cartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
		rc:          rc,
	}
}

func (c *cartService) AddCart(ctx context.Context, req entity.CartRequest) error {
	var cartModel models.Cart

	for _, product := range req.Products {
		p, err := c.productRepo.FindById(ctx, product.Id)
		if err != nil {
			return err
		}

		cartModel.Products = append(cartModel.Products, models.CartProduct{
			Product: models.Product{
				ID:    p.ID,
				Name:  p.Name,
				Price: p.Price,
			},
			Quantity: product.Qty,
		})
	}

	// get user id from context e.Request().Context()

	err := c.cartRepo.Store(ctx, cartModel)

	return err
}

func (c *cartService) GetCartById(ctx context.Context, id string) (*entity.Cart, error) {
	var cartEntity entity.Cart

	cart, err := c.cartRepo.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	cartEntity.ID = cart.ID.Hex()
	cartEntity.UserId = cart.CustomerID
	cartEntity.CreatedAt = cart.CreateAt
	cartEntity.UpdatedAt = cart.UpdateAt

	for _, product := range cart.Products {
		cartEntity.Products = append(cartEntity.Products, entity.CartProductResponse{
			Product: entity.Product{
				ID:    product.Product.ID.Hex(),
				Name:  product.Product.Name,
				Price: product.Product.Price,
			},
			Qty: product.Quantity,
		})
	}

	return &cartEntity, nil
}

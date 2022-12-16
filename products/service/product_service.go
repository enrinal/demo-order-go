package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/enrinal/demo-order-go/entity"

	"github.com/enrinal/demo-order-go/domain"
	"github.com/go-redis/redis/v8"
)

type productService struct {
	productRepository domain.ProductRepository
	rc                *redis.Client
}

func NewProductService(productRepository domain.ProductRepository, rc *redis.Client) domain.ProductService {
	return &productService{
		productRepository: productRepository,
		rc:                rc,
	}
}

func (p *productService) GetAll(ctx context.Context) ([]entity.Product, error) {
	products, err := p.productRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var productsEntity []entity.Product
	for _, product := range products {
		productsEntity = append(productsEntity, entity.Product{
			ID:    product.ID.Hex(),
			Name:  product.Name,
			Price: product.Price,
		})
	}

	return productsEntity, nil
}

func (p *productService) GetById(ctx context.Context, id string) (*entity.Product, error) {
	productCache, err := p.getProductFromRedis(ctx, id)
	if err == nil {
		return productCache, nil
	}

	product, err := p.productRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	productEntity := entity.Product{
		ID:    product.ID.Hex(),
		Name:  product.Name,
		Price: product.Price,
	}

	go func() {
		err := p.setProductToRedis(context.Background(), productEntity)
		if err != nil {
			log.Error().Err(err).Msg("error set product to redis")
		}
	}()

	return &productEntity, nil
}

func (p *productService) getProductFromRedis(ctx context.Context, id string) (*entity.Product, error) {
	var res *entity.Product
	op := p.rc.Get(ctx, entity.GetProductCacheKey(id))
	if op.Err() != nil {
		return nil, op.Err()
	}

	err := json.Unmarshal([]byte(op.Val()), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *productService) setProductToRedis(ctx context.Context, product entity.Product) error {
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}

	op := p.rc.Set(ctx, entity.GetProductCacheKey(product.ID), data, time.Duration(1)*time.Minute)
	return op.Err()
}

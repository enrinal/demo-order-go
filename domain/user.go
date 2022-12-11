package domain

import (
	"context"

	"github.com/enrinal/demo-order-go/entity"
	"github.com/enrinal/demo-order-go/models"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	Store(ctx context.Context, user models.User) error
}

type UserService interface {
	Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req entity.RegisterRequest) error
}

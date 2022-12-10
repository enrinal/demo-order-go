package service

import (
	"context"
	"github.com/enrinal/demo-order-go/constant"
	"github.com/enrinal/demo-order-go/entity"
	"github.com/enrinal/demo-order-go/models"
	"github.com/enrinal/demo-order-go/users/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req entity.RegisterRequest) error
}

type service struct {
	userRepo repository.UserRepo
}

func NewService(userRepo *repository.UserRepo) *service {
	return &service{*userRepo}
}

func (s *service) Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !checkPasswordHash(req.Password, user.Password) {
		return nil, nil
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = 3600

	t, err := token.SignedString([]byte(constant.SECRET))
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{
		Token: t,
	}, nil
}

func (s *service) Register(ctx context.Context, req entity.RegisterRequest) error {
	passHash, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = passHash

	err = s.userRepo.Store(ctx, models.User{
		Email:    req.Email,
		Password: req.Password,
	})

	return err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

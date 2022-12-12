package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/enrinal/demo-order-go/domain"

	"github.com/enrinal/demo-order-go/constant"
	"github.com/enrinal/demo-order-go/entity"
	"github.com/enrinal/demo-order-go/models"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo domain.UserRepository
	rc       *redis.Client
}

func NewService(userRepo domain.UserRepository, rc *redis.Client) domain.UserService {
	return &userService{
		userRepo: userRepo,
		rc:       rc,
	}
}

func (u *userService) Login(ctx context.Context, req entity.LoginRequest) (*entity.LoginResponse, error) {
	res, err := getCacheUserToken(u.rc, ctx, req.Email)
	if err == nil {
		log.Info().Msg("get user from cache")
		return &res, nil
	}

	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if !checkPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Duration(constant.ExpToken) * time.Second).Unix()

	t, err := token.SignedString([]byte(constant.Secret))
	if err != nil {
		return nil, err
	}

	// goroutine to set cache
	// fire and forget
	go func() {
		err := setCacheUserToken(u.rc, user.Email, entity.LoginResponse{Token: t})
		if err != nil {
			log.Error().Err(err).Msg("set cache user token")
		}
		log.Info().Msg("set cache user token success")
	}()

	return &entity.LoginResponse{
		Token: t,
	}, nil
}

func (u *userService) Register(ctx context.Context, req entity.RegisterRequest) error {
	// business logic
	passHash, err := hashPassword(req.Password)
	if err != nil {
		return err
	}

	req.Password = passHash

	// save to db
	err = u.userRepo.Store(ctx, models.User{
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

func cacheKey(email string) string {
	return fmt.Sprintf("user:%s", email)
}

func setCacheUserToken(rc *redis.Client, email string, res entity.LoginResponse) error {
	j, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return rc.Set(context.Background(), cacheKey(email), j, time.Duration(constant.ExpToken)*time.Second).Err()
}

func getCacheUserToken(rc *redis.Client, ctx context.Context, email string) (entity.LoginResponse, error) {
	var res entity.LoginResponse
	op := rc.Get(ctx, cacheKey(email))
	if op.Err() != nil {
		return res, op.Err()
	}

	err := json.Unmarshal([]byte(op.Val()), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

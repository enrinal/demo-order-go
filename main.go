package main

import (
	"github.com/enrinal/demo-order-go/config"
	productHdl "github.com/enrinal/demo-order-go/products/delivery"
	productRepo "github.com/enrinal/demo-order-go/products/repository"
	productSvc "github.com/enrinal/demo-order-go/products/service"
	userHdl "github.com/enrinal/demo-order-go/users/delivery"
	userRepo "github.com/enrinal/demo-order-go/users/repository"
	userSvc "github.com/enrinal/demo-order-go/users/service"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// redis
	redisHost := viper.GetString("redis.host")

	// connect to redis
	redisClient, err := config.ConnectRedis(redisHost)
	if err != nil {
		panic(err)
	}

	mongoUri := viper.GetString("mongo.uri")

	c, err := config.ConnectMongo(mongoUri)
	if err != nil {
		panic(err)
	}

	e := echo.New()

	userRepo := userRepo.NewUserRepo(c)
	userService := userSvc.NewService(userRepo, redisClient)
	userHdl.NewUserHandler(e, userService)

	productRepo := productRepo.NewProductRepo(c)
	productService := productSvc.NewProductService(productRepo, redisClient)
	productHdl.NewProductHandler(e, productService)

	e.Logger.Fatal(e.Start(":8080"))
}

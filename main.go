package main

import (
	"github.com/enrinal/demo-order-go/config"
	"github.com/enrinal/demo-order-go/users/delivery"
	"github.com/enrinal/demo-order-go/users/repository"
	"github.com/enrinal/demo-order-go/users/service"
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

	userRepo := repository.NewUserRepo(c)

	userService := service.NewService(userRepo, redisClient)

	delivery.NewUserHandler(e, userService)

	e.Logger.Fatal(e.Start(":8080"))
}

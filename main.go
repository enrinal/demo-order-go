package main

import (
	"os"

	"github.com/enrinal/demo-order-go/config"
	productHdl "github.com/enrinal/demo-order-go/products/delivery"
	productRepo "github.com/enrinal/demo-order-go/products/repository"
	productSvc "github.com/enrinal/demo-order-go/products/service"
	userHdl "github.com/enrinal/demo-order-go/users/delivery"
	userRepo "github.com/enrinal/demo-order-go/users/repository"
	userSvc "github.com/enrinal/demo-order-go/users/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	logger := zerolog.New(os.Stdout)
	logger.Level(zerolog.InfoLevel)

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

	newUserRepo := userRepo.NewUserRepo(c)
	userService := userSvc.NewService(newUserRepo, redisClient)
	userHandler := userHdl.NewUserHandler(userService)

	newProductRepo := productRepo.NewProductRepo(c)
	productService := productSvc.NewProductService(newProductRepo, redisClient)
	productHandler := productHdl.NewProductHandler(productService)

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	r := e.Group("/api/v1")
	// disable jwt middleware for register and login
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// product routes
	r.GET("/products", productHandler.GetAll)
	r.GET("/products/:id", productHandler.GetById)

	e.Logger.Fatal(e.Start(":8080"))
}

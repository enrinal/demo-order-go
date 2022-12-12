package main

import (
	"context"

	"github.com/enrinal/demo-order-go/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	mongoUri := viper.GetString("mongo.uri")

	c, err := config.ConnectMongo(mongoUri)
	if err != nil {
		panic(err)
	}

	// remove if exist
	_, err = c.Database("demo_order_go").Collection("orders").Indexes().DropOne(context.Background(), "status_1")

	// create index
	_, err = c.Database("demo_order_go").Collection("orders").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: map[string]int{"status": 1},
	})
	if err != nil {
		panic(err)
	}
}

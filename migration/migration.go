package main

import (
	"context"
	"github.com/enrinal/demo-order-go/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// create migration for mongodb include create index

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

	// create index
	_, err = c.Database("demo_order_go").Collection("users").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    map[string]int{"email": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		panic(err)
	}
}

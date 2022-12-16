package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/enrinal/demo-order-go/models"

	"github.com/enrinal/demo-order-go/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	//var products []models.Product

	// remove if exist
	_, err = c.Database("demo_order_go").Collection("products").Indexes().DropOne(context.Background(), "name_1")

	// create index unique
	_, err = c.Database("demo_order_go").Collection("products").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    map[string]int{"name": 1},
		Options: options.Index().SetUnique(true)})
	if err != nil {
		panic(err)
	}

	// drop collection
	err = c.Database("demo_order_go").Collection("products").Drop(context.Background())
	if err != nil {
		return
	}

	// add data
	var products []interface{}

	products = append(products, models.Product{
		ID:    primitive.NewObjectID(),
		Name:  "Apple",
		Price: 10000,
	})

	products = append(products, models.Product{
		ID:    primitive.NewObjectID(),
		Name:  "Orange",
		Price: 5000,
	})

	products = append(products, models.Product{
		ID:    primitive.NewObjectID(),
		Name:  "Banana",
		Price: 3000,
	})

	products = append(products, models.Product{
		ID:    primitive.NewObjectID(),
		Name:  "Mango",
		Price: 15000,
	})

	// insert data
	_, err = c.Database("demo_order_go").Collection("products").InsertMany(context.Background(), products)
	if err != nil {
		panic(err)
	}

}

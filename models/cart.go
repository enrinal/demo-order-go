package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// for mongodb

type CartProduct struct {
	Product
	Quantity int `json:"quantity" bson:"quantity"`
}

type Cart struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CustomerID string             `json:"customer_id" bson:"customer_id"`
	Products   []CartProduct      `json:"products" bson:"products"`
	CreateAt   string             `json:"create_at" bson:"create_at"`
	UpdateAt   string             `json:"update_at" bson:"update_at"`
}

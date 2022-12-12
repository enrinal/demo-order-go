package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// for mongodb

type Order struct {
	ID         primitive.ObjectID `bson:"_id"`
	CustomerID string             `bson:"customer_id"`
	CartID     string             `bson:"cart_id"`
	TotalPrice int                `bson:"total_price"`
	Status     string             `bson:"status"`
	CreateAt   string             `bson:"create_at"`
	UpdateAt   string             `bson:"update_at"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// for mongodb

type Product struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"name"`
	Price int                `bson:"price"`
}

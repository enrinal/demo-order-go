package models

// for mongodb

type Product struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Price int    `bson:"price"`
}

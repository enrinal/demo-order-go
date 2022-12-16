package repository

import (
	"context"
	"time"

	"github.com/enrinal/demo-order-go/constant"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/enrinal/demo-order-go/domain"
	"github.com/enrinal/demo-order-go/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type cartRepo struct {
	mgo *mongo.Client
}

func NewCartRepository(mgo *mongo.Client) domain.CartRepository {
	return &cartRepo{mgo}
}

func (c *cartRepo) Store(ctx context.Context, cart models.Cart) error {
	coll := c.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionCarts)

	cart.ID = primitive.NewObjectID()
	cart.CreateAt = time.Now().Format(time.RFC3339)
	cart.UpdateAt = time.Now().Format(time.RFC3339)

	_, err := coll.InsertOne(ctx, cart)
	if err != nil {
		return err
	}

	return nil
}

func (c *cartRepo) FindById(ctx context.Context, id string) (*models.Cart, error) {
	var cart *models.Cart

	coll := c.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionCarts)

	idCart, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = coll.FindOne(ctx, bson.M{
		"_id": idCart,
	}).Decode(&cart)
	if err != nil {
		return nil, err
	}

	return cart, nil
}

package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/enrinal/demo-order-go/constant"
	"github.com/enrinal/demo-order-go/domain"
	"github.com/enrinal/demo-order-go/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepo struct {
	mgo *mongo.Client
}

func NewProductRepo(mgo *mongo.Client) domain.ProductRepository {
	return &productRepo{
		mgo: mgo,
	}
}

func (p *productRepo) FindAll(ctx context.Context) ([]models.Product, error) {
	coll := p.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionProducts)
	var products []models.Product

	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (p *productRepo) FindById(ctx context.Context, id string) (*models.Product, error) {
	coll := p.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionProducts)
	var product models.Product

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = coll.FindOne(ctx, bson.M{
		"_id": objectId,
	}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

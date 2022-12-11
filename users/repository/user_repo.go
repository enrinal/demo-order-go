package repository

import (
	"context"
	"time"

	"github.com/enrinal/demo-order-go/domain"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/enrinal/demo-order-go/constant"
	"github.com/enrinal/demo-order-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	mgo *mongo.Client
}

func NewUserRepo(mgo *mongo.Client) domain.UserRepository {
	return &userRepo{
		mgo: mgo,
	}
}

func (u *userRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	// define collection
	collection := u.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionUsers)

	err := collection.FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) Store(ctx context.Context, user models.User) error {
	// define collection
	collection := u.mgo.Database(constant.MongoDatabaseName).Collection(constant.CollectionUsers)

	user.ID = primitive.NewObjectID() // generate new id in mongodb
	user.CreateAt = time.Now().String()
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
